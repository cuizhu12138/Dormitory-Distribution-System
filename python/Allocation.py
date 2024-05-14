from MPGA import Unit,Selection,Pool,Poolsys
import random as rd
from copy import deepcopy
from math import *
import sqlite3

dbpath='example.db'
roomMaxNum=0
sidlist=[]#待分寝的学生的学号列表
preSat={}

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
        for i in s.split(";"):
            tName,ans=i.split(":")
            ans=list(map(int,ans.split(",")))
            self[tName]=ans
            self['w'+tName]=1
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
    if sz<2:return None
    sat=0
    for i in range(0,sz):
        for j in range(i+1,sz):
            sat+=stuDataDict[sidlist[i]].satisfaction(stuDataDict[sidlist[j]])
            sat+=stuDataDict[sidlist[j]].satisfaction(stuDataDict[sidlist[i]])
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
            sat+=avgSatisfaction(self.dna[i:i+roomMaxNum])*roomMaxNum
            i+=roomMaxNum
        sat/=sz
        return sat
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
    roomMaxNum=6
    sidlist=[]#待分寝的学生的学号列表
    stuDataDict={}
    for i in range(1,60+1):
        sidlist.append(i)
        stuDataDict[i]=StudentData(i)
        stuDataDict[i].randInit()
def getStuData_db():
    global roomMaxNum,sidlist,stuDataDict
    roomMaxNum=6
    sidlist=[]
    stuDataDict={}
    with sqlite3.connect(dbpath) as db:
        cur=db.cursor()
        cur.execute("select * from UserHabbit")
        for uid,*anslist in cur:
            sidlist.append(uid)
            stu=StudentData(uid)
            stuDataDict[uid]=stu
            strans=[]
            for i,ans in enumerate(anslist):
                strans.append(f'{i}:{ans}')
            stu.strInit(';'.join(strans))
def writeAllocation(li:list[tuple]):
    with sqlite3.connect(dbpath) as db:
        cur=db.cursor()
        for (uid,roomid) in li:
            cur.execute(f"update DistributionResult set RoomNumber={roomid} where UID={uid}")
if __name__=='__main__':
    #getStuData_test()#getStuData_db()
    getStuData_db()
    ps=Poolsys(pi=0.5)
    ps.add(num=5,cls=Allocation,size=100,pc=0.9,pm=0.1,ps=10/100,select=Selection(0))
    # ps.decode()
    ps0=ps.copy()
    ps.run(50,end_bound=100)
    #ps.show()#outStuData
    ans=ps.decode()
    print(ans)
    writeAllocation(ans)
    
    # ps.decode()