// Hacker Earth
// Reverse a number

#include <iostream>

int main()
{
    int testcases;
    char n;
    std::cin>>testcases;
    while(testcases--)
    {
        std::cin>>n;
        char x = n.at(0);
        if(((int)x)%2)
            std::cout<<"ODD"<<std::endl;
        else
            std::cout<<"EVEN"<<std::endl;

    }
    return 0;
}
