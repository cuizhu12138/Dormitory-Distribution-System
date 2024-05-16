#基于SGA的多种群遗传算法MPGA
#需要实现Unit中除copy,__lt__以外的所有方法
#可以使用不同的Unit子类创建Pool,只要dna结构相同
#Pool遵从当前Unit子类的交叉变异等规则
import copy
import random as rd
from matplotlib import font_manager,ticker,pyplot as plt
import time
plt.rcParams['font.sans-serif'] = 'SimHei'#设置显示图片的字体
class Unit:
    def __init__(self):#初始化随机个体
        self.fitness=0
    def fit(self):#fitness适应度函数
        pass
    def cross(self,that):#corssover交叉算子->返回新个体或自身
        return self
    def muta(self):#mutation变异算子->返回新个体或自身
        return self
    def reverse(self):#进化逆转算子
        """
        重写此方法注意事项:
        1. self.fitness在调用此方法前已经更新(可以直接用)
        2. 方法返回的个体必须拥有正确的fitness成员变量(注意更新self.fitness)
        """
        return self
    def decode(self):#展示dna的解
        pass
    def __lt__(self,that):
        return self.fitness<that.fitness
    def copy(self):
        return copy.deepcopy(self)

class Selection:#锦标赛模拟线性排序选择函数,pmin/N为概率密度下限
    def __init__(self,pmin=0):
        self.p=pmin
    def __call__(self,pool):
        if rd.random()<self.p:
            return rd.choice(pool)
        u1,u2=rd.choice(pool),rd.choice(pool)
        return u1 if u1.fit()>u2.fit() else u2
    def __str__(self):
        return f'Selection({self.p})'

class Pool:#单个种群
    def __init__(self,cls:Unit,size=20,pc=0.9,pm=0.1,ps=0.05,select=Selection(),*,pattern={}):
        '''
        cls种群的个体类型
        size种群大小
        pc交叉率
        pm变异率
        ps精英保留率
        select选择函数
        pattern图形样式,plt.plot的字典参数,例:pattern={'c':'r'}
        '''
        self.pattern=pattern
        self.rule,self.size,self.pc,self.pm,self.ps,self.select=cls,size,pc,pm,ps,select
        self.axis_y_min,self.axis_y_max,self.axis_y_avg=[],[],[]#演化记录
        self.pool=[]#对象池
        for i in range(size):
            u=cls()
            u.fitness=u.fit()
            self.pool.append(u)
        self.log()
    def status(self):#统计种群状态
        'return (最小值,最大值,平均值,最劣解,最优解)'
        mn,mx,sm=min(self.pool),max(self.pool),0
        for u in self.pool:
            sm+=u.fitness
        return (mn.fitness,mx.fitness,sm/self.size,mn,mx)
    def log(self):#记录种群状态
        data=self.status()
        self.axis_y_min.append(data[0])
        self.axis_y_max.append(data[1])
        self.axis_y_avg.append(data[2])
        return data
    def next(self):
        saveNum=int(self.size*self.ps)
        newpool=self.pool[:saveNum]#保留精英
        for i in range(saveNum,self.size):
            son=copy.deepcopy(self.select(self.pool))
            if rd.random()<self.pc:#交叉
                son=self.rule.cross(son,self.select(self.pool))
            if rd.random()<self.pm:#突变
                son=self.rule.muta(son)
            son.fitness=self.rule.fit(son)
            son=son.reverse()#进化逆转
            newpool.append(son)
        self.pool=sorted(newpool,reverse=1)#每次迭代保持pool有序
        return self.log()
    def show(self):
        x=[i for i in range(len(self.axis_y_min))]
        plt.figure(figsize=(16,7), dpi=100)
        plt.plot(x,self.axis_y_min,c='r')
        plt.plot(x,self.axis_y_max,c='g')
        plt.plot(x,self.axis_y_avg,c='b')
        plt.show()
    def set(self,cls:Unit,size=20,pc=0.9,pm=0.1,ps=0.05,select=Selection(),*,pattern={}):#设置参数
        self.pattern=pattern
        self.rule,self.size,self.pc,self.pm,self.ps,self.select=cls,size,pc,pm,ps,select
    def settings(self):
        return {'cls':self.rule,'size':self.size,'pc':self.pc,'pm':self.pm,'ps':self.ps,'select':self.select,'pattern':self.pattern}
    def copy(self):
        return copy.deepcopy(self)
class Poolsys:#多种群管理
    def __init__(self,*arg,pi=0.5,**kwarg):
        self.pi=pi#移民率
        self.pools=[]
        self.poolmsg=[]
        self.elite_min=[]#精英种群历史
        self.elite_max=[]
        self.axis_x=[]
        if arg or kwarg:
            self.add(*arg,**kwarg)
    def add(self,num=1,*arg,**kwarg):
        for i in range(num):
            p=Pool(*arg,**kwarg)
            self.pools.append(p)
            self.poolmsg.append(self.pools[-1].status())#记录初始状态
    def log(self):#记录精英种群数据
        self.axis_x.append(self.axis_x[-1]+1)
        elitePool=list(zip(*self.poolmsg))[-1]
        self.elite_min.append(min(elitePool).fitness)
        self.elite_max.append(max(elitePool).fitness)
    def immigrate(self,pos1,pos2):#移民算子
        '移民方向pos1->pos2'
        if rd.random()<self.pi:
            self.pools[pos2].pool.append(self.poolmsg[pos1][-1])#pos2种群加入pos1最优个体
    def immigration(self):#移民操作
        for i in range(-1,len(self.pools)-1):
            self.immigrate(i,i+1)
    def run(self,rounds=500,end_bound=100):
        t_start=time.time()
        t_last=t_start
        if not self.axis_x:#初始化记录
            self.axis_x=[0]
            self.log()
            self.axis_x=[0]
        for round_num in range(1,rounds+1):
            t_now=time.time()
            print(f'r:{round_num}\tmax={self.elite_max[-1]}\ttime={t_now-t_last}')
            t_last=t_now
            
            self.poolmsg=[]
            for p in self.pools:
                self.poolmsg.append(p.next())
            self.immigration()
            self.log()
            if end_bound>0 and round_num>end_bound and self.elite_max[-end_bound]==self.elite_max[-1]:
                break#最大值保持end_bound轮
        totalTime=time.time()-t_start
        print(f'totalTime={totalTime} avgTime={totalTime/round_num}')
    def show(self,show_screen=True,outpath=None):
        plt.figure(figsize=(16,7),dpi=100)
        #plt.plot(self.axis_x,self.elite_max,c='r',label='best')
        for i,p in enumerate(self.pools):
            color,label=(rd.random(),rd.random(),rd.random()),f'population_{i}'
            kwarg=copy.deepcopy(p.pattern)
            if 'c' not in kwarg:
                kwarg['c']=color
            if 'label' not in p.pattern:
                kwarg['label']=label
            plt.plot(self.axis_x,p.axis_y_max,**kwarg)
            plt.text(self.axis_x[-1],p.axis_y_max[-1],f'{i}',c=color)
        plt.legend(loc='best')
        if show_screen:
            plt.show()
        if outpath:
            plt.savefig(outpath)
    def draw(self,c,label='None'):#绘制所有曲线,同色
        plt.plot(self.axis_x,self.elite_max,c=c,label=label)
        for i,p in enumerate(self.pools):
            plt.plot(self.axis_x,p.axis_y_max,c=c)
        plt.legend(loc='best')
    def decode(self):
        return max(list(zip(*self.poolmsg))[-1]).decode()
    def copy(self):
        return copy.deepcopy(self)
