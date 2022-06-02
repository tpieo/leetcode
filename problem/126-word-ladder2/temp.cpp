#include <string>
#include <vector>
#include <unordered_map>
#include <unordered_set>
using namespace std;

void backtracking(string& src, string& dst, unordered_map<string, vector<string>>& next, vector<vector<string>>& ans, vector<string> path) {
	if (src == dst) {
		ans.push_back(path);
		return;
	}
	for (auto& word : next[src]) {
		path.push_back(word);
		backtracking(word, dst, next, ans, path);
		path.pop_back();
	}
}

bool isOneLetter(string word1, string word2) {
	int c = 0;
	for (size_t i = 0; i < word1.size(); i++ ) {
		if (word1[i] != word2[i]) {
			c++;
		}
	}
	if (c == 1) {
		return true;
	}
	return false;
}

class Solution {
public:
    vector<vector<string>> findLadders(string beginWord, string endWord, vector<string>& wordList) {
        unordered_set<string> wordSet;
        for (auto& word : wordList) {
            wordSet.insert(word);
        }
        vector<vector<string>> ans;
        if (wordSet.find(endWord) == wordSet.end()) {
            return ans;
        }
        wordSet.erase(beginWord);
        unordered_map<string, vector<string>> next;
        unordered_set<string> q1{beginWord};
        bool found = false;
        for (;!q1.empty();) {
            unordered_set<string> q;
            for (auto& word1 : q1) {
                for (auto& word2 : wordSet) {
                	if (isOneLetter(word1, word2)) {
                		next[word1].push_back(word2);
                		q.insert(word2);
                		if (word2 == endWord) {
                			found = true;
                		}
                	}
                }
            }
            if (found) {
            	break;
            }
            for (auto& word : q) {
            	wordSet.erase(word);
            }
            q1 = q;
        }
        if (found) {
        	vector<string> path{beginWord};
        	backtracking(beginWord, endWord, next, ans, path);
        }
        return ans;
    }
};

int main()
{
    string beginWord = "hit";
    string endWord = "cog";
    vector<string> wordList = {"hot","dot","dog","lot","log","cog"};
    vector<vector<string>> ans = Solution().findLadders(beginWord, endWord, wordList);
}