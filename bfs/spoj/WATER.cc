#include <cstdio>
#include <cstring>
#include <iostream>
#include <queue>
using namespace std;

#define ROWS 100
#define COLS 100

using P = pair<int, pair<int, int>>; // value, [rowIndex, colIndex]
using PQ = priority_queue<P, vector<P>, greater<P>>;

int rows, cols;
bool visited[ROWS][COLS];
int g[ROWS][COLS];

P make_P(int value, int row, int col) {
  return make_pair(value, make_pair(row, col));
}

void clear() {
  memset(visited, 0, sizeof visited);
  memset(g, 0, sizeof g);
}

int rowOff[] = {-1, 0, 0, +1};
int colOff[] = {0, +1, -1, 0};

bool inBound(int newRow, int newCol) {
  return (newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols);
}

void floodFill(PQ &q, int x, int y, int startingHeight, int *count) {
  // cout << x << " " << y << " " << g[x][y] << "\n";
  if (visited[x][y]) {
    return;
  }
  *count += max(0, startingHeight - g[x][y]);
  visited[x][y] = true;
  for (int i = 0; i < 4; i++) {
    int newRow = x + rowOff[i];
    int newCol = y + colOff[i];
    if (inBound(newRow, newCol)) {
      if (g[newRow][newCol] >= startingHeight) {
        q.push(make_P(g[newRow][newCol], newRow, newCol));
      } else {
        floodFill(q, newRow, newCol, startingHeight, count);
      }
    }
  }
}

int bfs(PQ &q) {
  int total = 0;
  while (!q.empty()) {
    P curr = q.top();
    q.pop();
    int x = curr.second.first;
    int y = curr.second.second;
    // printf("visiting: (%d, %d) = %d (%d)\n", x, y, g[x][y], curr.first);
    if (visited[x][y]) {
      continue;
    }
    int t = 0;
    floodFill(q, x, y, curr.first, &t);
    total += t;
  }
  return total;
}

void solve() {
  cin >> rows >> cols;
  for (int r = 0; r < rows; r++) {
    for (int c = 0; c < cols; c++) {
      cin >> g[r][c];
    }
  }
  PQ q;
  for (int i = 0; i < cols; i++) {
    q.push(make_P(g[0][i], 0, i));
    q.push(make_P(g[rows - 1][i], rows - 1, i));
  }
  for (int i = 0; i < rows; i++) {
    q.push(make_P(g[i][0], i, 0));
    q.push(make_P(g[i][cols - 1], i, cols - 1));
  }
  // make the corner visited
  visited[0][0] = true;
  visited[0][cols - 1] = true;
  visited[rows - 1][0] = true;
  visited[rows - 1][cols - 1] = true;
  cout << bfs(q) << "\n";
}

int main() {
  int T;
  cin >> T;
  for (int i = 0; i < T; i++) {
    clear();
    solve();
  }
}