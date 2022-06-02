#include <vector>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <queue>
using namespace std;

void backtracking(string src, string dst, unordered_map<string, vector<string>>& next, vector<vector<string>>& ans, vector<string> path) {
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

class Solution {
public:
    vector<vector<string>> findLadders(string beginWord, string endWord, vector<string>& wordList) {
    	vector<vector<string>> ans;
    	unordered_set<string> wordSet;
    	for (auto& word : wordList) {
    		wordSet.insert(word);
    	}
    	if (!wordSet.count(endWord)) {
    		return ans;
    	}
    	unordered_map<string, vector<string>> next;
    	wordSet.erase(beginWord);
    	wordSet.erase(endWord);
    	unordered_set<string> q1{beginWord};
    	bool found = false;
    	while (!q1.empty()) {
    		unordered_set<string> q;
    		for (auto& word : q1) {
    			for (size_t i = 0; i < word.size(); i++) {
    				string tempWord = word;
    				for (int j = 0; j < 26; j++) {
    					tempWord[i] = 'a' + j;
    					if (wordSet.find(tempWord) != wordSet.end()) {
    						q.insert(tempWord);
    						next[word].push_back(tempWord);
    					}
    					if (tempWord == endWord) {
							next[word].push_back(endWord);
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