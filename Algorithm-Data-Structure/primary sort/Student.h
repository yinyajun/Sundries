#include <iostream>
#include <string>

using namespace std;

struct Student{
    string name;
    int score;

    bool operator<(const Student &otherStudent){
        return score<otherStudent.score;
    }

    //要想对类的内容输出，必须重载<<,为了方便读取对象内部的成员，设置为友元函数
    //为了达到连续输出的效果，设置返回类型为引用类型,os<<'da'返回的还是ostream对象
    //第一个参数为输出流对象，第二个参数为输出对象（使用引用类型防止产生临时对象；使用常量的引用，防止修改实参） 
    friend ostream& operator<<(ostream& os, const Student& student){
        os<<"Student:"<<student.name<<" "<<student.score<<endl;
        return os;
    }
};

