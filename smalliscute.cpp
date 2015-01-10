// HackerEarth - Practice
// Small is cute

#include <iostream>
#include <set>

int main()
{
    unsigned int n,a,b;
    std::set<unsigned int> divisibles;
    std::cin>>n>>a>>b;

    for(unsigned int i=0;i<n;++i)
    {
        if(i%a==0)
            divisibles.insert(i);
        else if(i>=b && i%b==0)
            divisibles.insert(i);
    }
    long int sum=0;
    for(std::set<unsigned int>::iterator i=divisibles.begin();i!=divisibles.end();i++)
        sum+=*i;
    std::cout<<sum;
    return 0;
}
