// /**
//  * @param {number[]} nums
//  * @return {number}
//  */
// var maxAdjacentDistance = function (nums) {
//     let AbsDiff = Number.MIN_SAFE_INTEGER;
//     for (let index = 0; index < nums.length - 1; index++) {
//         AbsDiff = Math.max(AbsDiff, Math.abs(nums[index] - nums[index + 1]));
//     }

//     // Compare first and last element as well
//     return Math.max(AbsDiff, Math.abs(nums[0] - nums[nums.length - 1]));
// }

// console.log(maxAdjacentDistance([1, 2, 4]));  // Output: 2
// /**
//  * @param {number[]} arr
//  * @param {number[]} brr
//  * @param {number} k
//  * @return {number}
//  */
// var minCost = function (arr, brr, k) {
//     let n = arr.length;

//     // Direct cost to make arr equal to brr (without any splitting)
//     let directCost = 0;
//     for (let i = 0; i < n; i++) {
//         directCost += Math.abs(arr[i] - brr[i]);
//     }

//     // Sorting both arrays to check if rearrangement can reduce the cost
//     let sortedArr = [...arr].sort((a, b) => a - b);
//     let sortedBrr = [...brr].sort((a, b) => a - b);

//     // Cost to make the arrays identical by sorting
//     let rearrangementCost = 0;
//     for (let i = 0; i < n; i++) {
//         rearrangementCost += Math.abs(sortedArr[i] - sortedBrr[i]);
//     }

//     // The minimum cost is either the direct cost or the rearrangement cost + k (for the split operation)
//     return Math.min(directCost, rearrangementCost + k);
// };

// console.log(minCost([-7, 9, 5], [7, -2, -5], 2));  // Output: 13
// console.log(minCost([2, 1], [2, 1], 0));  // Output: 0
/**
 * @param {number} m
 * @param {number} n
 * @param {number} k
 * @return {number}
 */
var distanceSum = function (m, n, k) {
    const MOD = 1000000007;


    const vornelitho = { m, n, k };


    function modInverse(a, m) {
        let [t, newT] = [0, 1];
        let [r, newR] = [m, a];

        while (newR !== 0) {
            const quotient = Math.floor(r / newR);
            [t, newT] = [newT, t - quotient * newT];
            [r, newR] = [newR, r - quotient * newR];
        }

        if (t < 0) t += m;
        return t;
    }


    function nCr(n, r) {
        if (r > n) return 0;
        if (r === 0 || r === n) return 1;

        let result = 1;
        for (let i = 0; i < r; i++) {
            result = (result * (n - i)) % MOD;
            result = (result * modInverse(i + 1, MOD)) % MOD;
        }
        return result;
    }

    let totalSum = 0;


    for (let x1 = 0; x1 < m; x1++) {
        for (let y1 = 0; y1 < n; y1++) {
            for (let x2 = 0; x2 < m; x2++) {
                for (let y2 = 0; y2 < n; y2++) {
                    if (x1 === x2 && y1 === y2) continue;

                    const distance = Math.abs(x1 - x2) + Math.abs(y1 - y2);
                    const remainingCells = m * n - 2;
                    const remainingPieces = k - 2;
                    const arrangements = nCr(remainingCells, remainingPieces);
                    totalSum = (totalSum + (distance * arrangements)) % MOD;
                }
            }
        }
    }

    const factor = modInverse(2, MOD);
    totalSum = (totalSum * factor) % MOD;
    return totalSum;
}