public class Solution {
	/*
     * @param string: An array of Char
     * @param length: The true length of the string
     * @return: The true length of new string
     */
	public int replaceBlank(char[] string, int length) {
        if (length==0){
            return 0;
        }
		for (int i=0; i<length;) {
			if (string[i] == ' '){
				replace(string,i,3-1);
                i += 3;
                length += 2;
			} else {
                i++;
            }
        }
		return length;
    }
    
	public void replace(char[] string,int pos,int num){
		for (int i=string.length-1;i>=pos&&i>=0&&i-num>=0;i--){
			string[i]=string[i-num];
        }
        string[pos]='%';
		string[pos+1]='2';
		string[pos+2]='0';
    }

    public static void main(String[] args){
        Solution solution = new Solution();
        char[] example1 = new char[20];
        char[] array1 = "Mr John Smith".toCharArray();
        System.arraycopy(array1,0,example1,0,array1.length);
      
        char[] example2 = new char[20];
        char[] array2 = " helloworld".toCharArray();
        System.arraycopy(array2,0,example2,0,array2.length);

        char[] example3 = new char[20];
        char[] array3 = "   ".toCharArray();
        System.arraycopy(array3,0,example3,0,array3.length);
        
		System.out.println(solution.replaceBlank(example1, 13));
        System.out.println(solution.replaceBlank(example2, 11));
        System.out.println(solution.replaceBlank(example3, 3));
        System.out.println(solution.replaceBlank(null, 0));
        
	}
}