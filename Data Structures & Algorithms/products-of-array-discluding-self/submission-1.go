func productExceptSelf(nums []int) []int {
    n := len(nums)
    products := make([]int, n)

    prefix := 1
    for i := 0; i < n; i++ {
        products[i] = prefix
        prefix *= nums[i]
    }

    suffix := 1
    for i := n - 1; i >= 0; i-- {
        products[i] *= suffix
        suffix *= nums[i]
    }

    return products
}