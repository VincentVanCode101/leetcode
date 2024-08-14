import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

class Solution {
    public List<String> fizzBuzz(int n) {
        Map<Integer, String> fizzBuzzMap = new HashMap<>();
        fizzBuzzMap.put(3, "Fizz");
        fizzBuzzMap.put(5, "Buzz");

        List<String> answerList = new ArrayList<>();
        for (int i = 1; i <= n; i++) {
            StringBuilder sb = new StringBuilder();
            for (Map.Entry<Integer, String> entry : fizzBuzzMap.entrySet()) {
                if (i % entry.getKey() == 0) {
                    sb.append(entry.getValue());
                }
            }
            if (sb.length() == 0) {
                sb.append(i);
            }
            answerList.add(sb.toString());
        }
        return answerList;
    }
}
