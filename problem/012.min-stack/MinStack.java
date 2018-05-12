import java.util.Stack;

public class MinStack {
    private Stack<Integer> stack;
    private Stack<Integer> minStack;

    public MinStack() {
        stack = new Stack<>();
        minStack = new Stack<>();
    }

    /*
     * @param number: An integer
     * 
     * @return: nothing
     */
    public void push(int number) {
        stack.push(number);
        if (minStack.empty() || number <= minStack.peek()) {
            minStack.push(number);
        }
    }

    /*
     * @return: An integer
     */
    public int pop() {
        int num = stack.pop();
        if (num == minStack.peek()) {
            minStack.pop();
        }
        return num;
    }

    /*
     * @return: An integer
     */
    public int min() {
        return minStack.peek();
    }

    public static void main(String[] args) {
        MinStack stack = new MinStack();
        stack.push(1);
        stack.push(1);
        stack.push(1);
        System.out.println(stack.min());
        System.out.println(stack.pop());
        System.out.println(stack.min());
        System.out.println(stack.pop());

    }
}