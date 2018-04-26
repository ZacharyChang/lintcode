import java.util.Stack;

public class MyQueue {

    private Stack<Integer> stack1;
    private Stack<Integer> stack2;

    public MyQueue() {
        // do intialization if necessary
        stack1 = new Stack<>();
        stack2 = new Stack<>();
    }

    /*
     * @param element: An integer
     * @return: nothing
     */
    public void push(int element) {
        stack1.push(element);
        stack2.push(element);
    }

    /*
     * @return: An integer
     */
    public int pop() {
        stack1.clear();
        while(!stack2.empty()){
            stack1.push(stack2.pop());
        }
        int result = stack1.pop();
        stack2.clear();
        while(!stack1.empty()){
            stack2.push(stack1.pop());
        }
        stack1.clear();
        stack1.addAll(stack2);
        return result;
    }

    /*
     * @return: An integer
     */
    public int top() {
        stack1.clear();
        while(!stack2.empty()){
            stack1.push(stack2.pop());
        }
        int result = stack1.peek();
        stack2.clear();
        while(!stack1.empty()){
            stack2.push(stack1.pop());
        }
        stack1.clear();
        stack1.addAll(stack2);
        return result;
    }

    private void switchStack() {
        while (!stack2.empty()) {
            stack1.push(stack2.pop());
        }
    }

    public static void main(String[] args){
        MyQueue queue = new MyQueue();
        queue.push(1);
        System.out.println(queue.pop() == 1);
        queue.push(2);
        queue.push(3);
        System.out.println(queue.top() == 2);
        System.out.println(queue.pop() == 2);
	}
}