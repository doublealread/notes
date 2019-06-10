# _*_coding:utf-8 _*_
# Author:Oldyuan
# ChangeTime: 2019/6/6 11:11
# FileName: 1-面向过程的编程.py


import json

def intersctive():
    name = input('>>>:').strip()
    pwd = input(">>>:").strip()
    return {
        'name':name,
        'pwd':pwd
    }


def check(user_info):
    is_valid = True

    if len(user_info['name']) == 0:
        print('空')

        is_valid = False

    if len(user_info['pwd']) < 6:
        print('小')

        is_valid = False

    return {
        'is_valid:':is_valid,
        'user_info:':user_info
    }

def register(check_info):
    # print(check_info['is_valid'])
    # if check_info.setdefault('is_valid',0):
    if check_info['is_valid']:
        with open('db.json','w',encoding='utf-8') as f:
            json.dump(check_info['user_info'],f)





if __name__ == '__main__':
    user_info = intersctive()
    check_info = check(user_info)
    print(check_info)
    register(check_info)




