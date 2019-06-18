# 面向对象编程


* 改
```python
stu1.Name = "lierya"
print(stu1.__dict__)
```
* 删除
```python
    del stu1.Name
    print(stu1.__dict__)
```
* 增加
```python
stu1.class_name = 'python开发'
print(stu1.__dict__)
```

### 属性的查找与方法的绑定

```python
class People:
    school = "NIUJING"
    def __init__(self,name,age,sex):
        self.Name = name
        self.Age = age
        self.Sex = sex
    
    def learn(self):
        print('learn，%s',self.Name)
    def eat(self):
        print('eat')
        
stu1 = People('wangerya','nv',38)
stu2 = People('lisanpao','nan',40)

# 对象： 特征与技能的结合体
#类：类是一些列对象相似的特征与相似的技能的结合体

#类中的数据属性：是所有对象共有的
print(stu1.school.id(stu1.school))
#类中的函数属性:是绑定到不同的对象是不同的绑定方法,对象绑定方式时，会吧对象本身当做第一个
#       传入，传给self
People.learn(stu1)
#输出结果： learn ，wangerya
print(stu1.learn())
stu1.x = "from stu1"
People.x = 'from Persion class'

#首先会查找对象中的参数，如果对象中没有，才会查找类中的参数。
print(stu1.x)

```
### python一切皆对象

















