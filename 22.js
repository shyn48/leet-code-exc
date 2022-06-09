/**
 * @param {number} n
 * @return {string[]}
 */

var generateParenthesis = function(n) {
    const stack = []
    const result = []

    const backtrack = (opened, closed) => {
        if (opened === n && closed === n) {
            result.push(stack.join(""))
            return 
        }
        
        if (opened < n) {
            stack.push("(")
            backtrack(opened  + 1, closed)
            stack.pop()
        }
        
        if (closed < opened) {
            stack.push(")")
            backtrack(opened, closed + 1)
            stack.pop()
        }
    }
    
    
    backtrack(0,0)
    
    return result
};
