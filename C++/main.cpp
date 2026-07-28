#include <iostream>

int main(){
    auto result = (10 <=> 20) > 0;
    std::count << result << std::endl;
}