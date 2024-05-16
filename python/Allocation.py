import os
from MPGA import Unit,Selection,Pool,Poolsys
import random as rd
from copy import deepcopy
from math import *
import sqlite3
import argparse

parser=argparse.ArgumentParser()
parser.add_argument('--second',action="store_true",help="启用二次分配")
parser.add_argument('--dbpath',default="example.db",metavar="path",action="store",help="sqlite数据库文件路径")
parser.add_argument('--show',action="store_true",help="启用选项将输出适应度变化图到屏幕")
parser.add_argument('--image',metavar="path",help="输出适应度变化图")
parser.add_argument('--room_size',default=6,type=int,metavar="num",help="每个寝室的大小")
parser.add_argument('--rounds',default=50,type=int,metavar="num",help="迭代次数")
parser.add_argument('--end_bound',default=100,type=int,metavar="num",help="终止条件(最近n次迭代的适应度不再变化)")
parser.add_argument('--pi',default=0.5,type=float,help="移民率")
parser.add_argument("--num",default=5,type=int,help="种群数量")
parser.add_argument("--size",default=100,type=int,help="种群大小(种群内个体数量)")
parser.add_argument("--pc",default=0.9,type=float,help="交叉率")
parser.add_argument("--pm",default=0.1,type=float,help="变异率")
parser.add_argument("--ps",default=10/100,type=float,help="精英保留率")
sysarg=parser.parse_args()

dbpath=sysarg.dbpath
roomMaxNum=sysarg.room_size
sidlist=[]#待分寝的学生的学号列表
preSat={}

tot_weight=0

class StudentData(dict):
    def __init__(self,sid):#获取问卷信息,权重
        super().__init__()
        self.sid=sid
        #self['q1']=0~3 self['w1']=0.xx
    def randInit(self):
        for i in range(6,33+1):
            self[f'q{i}']=[rd.randint(0,3)]
            self[f'wq{i}']=1
            # self['q24']=[]
    def strInit(self,s):
        global tot_weight
        tot_weight=0
        for i in s.split(";"):
            tName,ans=i.split(":")
            ans=list(map(int,ans.split(",")))
            self[tName]=ans
            self['w'+tName]=1
            tot_weight+=1
    def satisfaction(self,that):#仅自己对对方的评价
        sat=preSat.get((self.sid,that.sid),0)
        if sat:return sat
        for tName,li in self.items():
            if tName[0]=='w':continue
            if len(li)==1:
                val=li[0]
                sat+=self['w'+tName]*(1/(1+(that[tName][0]-val)**2))
        preSat[(self.sid,that.sid)]=sat
        return sat
        #---old---
        #1.基本信息
        pass
        #2.生活作息
        if self['q8']==0:
            sat+=self['w6']*(1/(1+(that['q6']-self['q6'])**2))
            sat+=self['w7']*(1/(1+(that['q7']-self['q7'])**2))
        else:
            sat+=self['w6']
            sat+=self['w7']
        #特判爱好匹配 self['q24']是列表
        if self['q24']:
            sameNum=0
            for i in self['q24']:
                if i in that['q24']:
                    sameNum+=1
            sat+=self['w24']*(sameNum/len(self['q24']))
        #方差倒数匹配(默认)
        for i in [*range(9,23+1),*range(25,33+1)]:
            sat+=self[f'w{i}']*(1/(1+(that[f'q{i}']-self[f'q{i}'])**2))
        #存表
        preSat[(self.sid,that.sid)]=sat
        return sat
stuDataDict={}
def avgSatisfaction(sidlist:list):#将列表中学生两两匹配,计算平均满意度
    sz=len(sidlist)
    if sz<1:return None
    sat=0
    for i in range(0,sz):
        for j in range(i+1,sz):
            sat+=stuDataDict[sidlist[i]].satisfaction(stuDataDict[sidlist[j]])
            sat+=stuDataDict[sidlist[j]].satisfaction(stuDataDict[sidlist[i]])
    if sz==1:return stuDataDict[sidlist[0]].satisfaction(stuDataDict[sidlist[0]])
    return sat/sz/(sz-1)#取均值

class Allocation(Unit):
    def __init__(self):
        super().__init__()
        self.dna=deepcopy(sidlist)
        rd.shuffle(self.dna)
    def fit(self):
        i=0
        sz=len(self.dna)
        sat=0
        while i<sz:
            room_stulist=self.dna[i:i+roomMaxNum]
            sat+=avgSatisfaction(room_stulist)*len(room_stulist)
            i+=roomMaxNum
        sat/=sz
        return sat/tot_weight
    def muta(self):
        sz=len(self.dna)
        i,j=rd.randint(0,sz-1),rd.randint(0,sz-1)
        self.dna[i],self.dna[j]=self.dna[j],self.dna[i]
        return self
    def cross(self,that):
        sz=len(self.dna)
        pos=rd.randint(0,sz-1)
        #print(f'pos={pos}')
        newdna=[0]*sz
        st=set()
        for i in range(pos,sz):
            newdna[i]=that.dna[i]
            st.add(that.dna[i])
        j=0
        for i in range(pos):
            while self.dna[j] in st:
                j+=1
            newdna[i]=self.dna[j]
            j+=1
        self.dna=newdna
        return self
    def reverse(self):
        # return self
        old_fitness=self.fitness
        sz=len(self.dna)
        st,ed=rd.randint(0,sz-1),rd.randint(0,sz-1)
        if st>ed:st,ed=ed,st

        for i in range((ed-st+1)//2):
            self.dna[st+i],self.dna[ed-i]=self.dna[ed-i],self.dna[st+i]
        new_fitness=self.fit()

        if new_fitness>=old_fitness:
            #采纳逆转
            self.fitness=new_fitness
            return self
        else:
            #还原dna
            for i in range((ed-st+1)//2):
                self.dna[st+i],self.dna[ed-i]=self.dna[ed-i],self.dna[st+i]
            return self
    def decode(self):#返回list[(学生id,对应寝室号)]
        ans=[]
        for pos,id in enumerate(self.dna):
            ans.append((id,pos//roomMaxNum))
        return ans
def getStuData_test():
    global roomMaxNum,sidlist,stuDataDict
    #初始化,获取学生信息和寝室人数
    # roomMaxNum=6
    sidlist=[]#待分寝的学生的学号列表
    stuDataDict={}
    for i in range(1,60+1):
        sidlist.append(i)
        stuDataDict[i]=StudentData(i)
        stuDataDict[i].randInit()
def getStuData_db():
    global roomMaxNum,sidlist,stuDataDict
    # roomMaxNum=6
    sidlist=[]
    stuDataDict={}
    with sqlite3.connect(dbpath) as db:
        cur=db.cursor()
        cur.execute("select * from user_questionnaire_data")#UserHabbit
        for uid,*anslist in cur:
            sidlist.append(uid)
            stu=StudentData(uid)
            stuDataDict[uid]=stu
            strans=[]
            for i,ans in enumerate(anslist):
                strans.append(f'{i}:{ans}')
            stu.strInit(';'.join(strans))
def writeAllocation(li:list[tuple],second=False):
    "Args:\n\tli: list[tuple[uid,roomid]]"
    with sqlite3.connect(dbpath) as db:
        cur=db.cursor()
        for (uid,roomid) in li:
            if second:
                cur.execute(f"update DistributionResult set ReassignResult={roomid} where UID={uid}")
            else:
                cur.execute(f"update DistributionResult set RoomNumber={roomid} where UID={uid}")
def ans_merge(*ansli):#合并
    sidlist=[]#
    cnt=0
    for oldans in ansli:
        mp=dict()#oldrid->newrid
        for sid,rid in oldans:
            if rid not in mp:
                mp[rid]=cnt
                cnt+=1
            sidlist.append((sid,mp[rid]))
    return sidlist
def main_mpga_base(sids:list,title_name='',*,file_count=[0]):#不分男女
    # ps=Poolsys(pi=0.5)
    # ps.add(num=5,cls=Allocation,size=100,pc=0.9,pm=0.1,ps=10/100,select=Selection(0))
    global sidlist
    sidlist=sids
    ps=Poolsys(pi=sysarg.pi)
    ps.add(num=sysarg.num,cls=Allocation,size=sysarg.size,pc=sysarg.pc,pm=sysarg.pm,ps=sysarg.ps,select=Selection(0))
    # ps.decode()
    ps0=ps.copy()
    ps.run(sysarg.rounds,end_bound=sysarg.end_bound)
    #ps.show()#outStuData
    path=None
    if sysarg.image:
        paths=os.path.splitext(sysarg.image)
        path=paths[0]+f'_{file_count[0]}'+paths[1]
        file_count[0]+=1
    if sysarg.show or sysarg.image:
        ps.show(sysarg.show,path,title_name)
    return ps.decode()
def main_mpga(sids:list):#分男女寝
    sex=dict()
    with sqlite3.connect(dbpath) as db:
        cur=db.cursor()
        cur.execute("select uid,sex from user_base_infos")
        for uid,stusex in cur:
            sex[uid]=int(stusex)
    sids0=[]
    sids1=[]
    for sid in sids:
        if sex[sid]==0:
            sids0.append(sid)
        else:
            sids1.append(sid)
    ans1=main_mpga_base(sids0,'男')#男
    ans2=main_mpga_base(sids1,'女')#女
    return ans_merge(ans1,ans2)
    
def main_first():#初次分配
    #getStuData_test()#getStuData_db()
    getStuData_db()
    ans=main_mpga(sidlist)
    print(ans)
    writeAllocation(ans)
def main_second():#二次分配
    global sidlist
    getStuData_db()
    rids=set()#需要重分配的房间id
    allo={}#当前分配方案
    with sqlite3.connect(dbpath) as db:
        cur=db.cursor()
        cur.execute(f"select UID,RoomNumber,DecisionForReassign from DistributionResult")
        for uid,rid,reAllocation in cur:
            allo[uid]=rid
            if reAllocation:rids.add(rid)
    sidlist=[]#重置分寝列表
    for uid,rid in allo.items():
        if rid in rids:
            sidlist.append(uid)
    
    if rids:#需要重分配
        ans=main_mpga(sidlist)
        first_rid=list(rids)
        for uid,rid in ans:#映射原本的房间
            allo[uid]=first_rid[rid]
    
    writeAllocation(list(allo.items()),True)

def main():
    if sysarg.second:
        main_second()
        return
    main_first()

if __name__=='__main__':
    main()