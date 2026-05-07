def dnf(arr)
  low = 0
  mid = 0
  high = arr.length - 1

  while mid <= high
    if arr[mid] == 0
      arr[low], arr[mid] = arr[mid], arr[low]
      low = low + 1
      mid = mid + 1
    elsif arr[mid] == 1
      mid = mid + 1
    else
      arr[mid], arr[high] = arr[high], arr[mid]
      high = high - 1
    end
  end
  return arr
end

p(dnf([1,0,0,1,2,1,0]))
