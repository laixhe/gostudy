#include <iostream>
#include "Singleton.h"

// 单例模式

Singleton::Singleton() {
    std::cout<<"Singleton()"<<std::endl;
}

Singleton::~Singleton() {
    std::cout<<"~Singleton()"<<std::endl;
}

// 利用局部静态变量只会初始化一次，而且还是线程安全的
// 注意返回的是引用
// C++11 之前不能这么写
Singleton& Singleton::getInstance() {
    // 局部静态变量
    static Singleton instance;
    return instance;
}
