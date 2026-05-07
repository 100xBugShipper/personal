function factorial(n)
    if n < 2 then
        return 1
    else
        return n * factorial(n-1)
    end
end

print(factorial(5))
