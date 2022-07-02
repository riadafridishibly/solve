#include <iostream>
#include <queue>
#include <cstring>

#define N 8

using namespace std;
using P = pair<int, int>;
bool visited[N][N];
int depth[N][N];

void clear()
{
    memset(visited, 0, sizeof visited);
    memset(depth, 0, sizeof depth);
}

int x[] = {-2, -2, -1, -1, +1, +1, +2, +2};
int y[] = {+1, -1, +2, -2, +2, -2, +1, -1};

P convertToPair(string s)
{
    return make_pair(s[0] - 'a', s[1] - '0' - 1);
}

int bfs(P source, P dest)
{
    // printPair(source);
    // printPair(dest);
    queue<P> q;
    q.push(source);
    int lvl = 0;
    bool done = false;
    while (!q.empty() && !done)
    {
        int sz = q.size();
        for (int j = 0; j < sz; j++)
        {
            P curr = q.front();
            q.pop();
            if (curr == dest)
            {
                done = true;
                break;
            }
            if (visited[curr.first][curr.second])
            {
                continue;
            }
            visited[curr.first][curr.second] = true;
            // Generate
            for (int i = 0; i < 8; i++)
            {
                int newX = curr.first + x[i];
                int newY = curr.second + y[i];
                if (newX >= 0 && newX < 8 && newY >= 0 && newY < 8)
                {
                    depth[newX][newY] = depth[curr.first][curr.second] + 1;
                    q.emplace(newX, newY);
                }
            }
        }
        lvl++;
    }
    // cout << "level: " << --lvl << "\n";
    return depth[dest.first][dest.second];
}

void solve()
{
    string start, end;
    cin >> start >> end;
    cout << bfs(convertToPair(start), convertToPair(end)) << "\n";
}

int main()
{
    int T;
    cin >> T;
    for (int i = 0; i < T; i++)
    {
        clear();
        solve();
    }
}