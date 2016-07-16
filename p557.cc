#include <iostream>
#include <set>

using std::cout;
using std::endl;
using std::set;

// a is a vertex at (0, 0)
// b is a vertex at (0, area)
// c is a vertex at (2, 0)
// d is the point between a and b
// e is a point between a and c
// f is the intersection of bd and ce
//
// Algorithm is to check if bdf is integral.
// The other areas will be integral if bdf is by construction. 
long long findIntegralAreas(int area) {
  long long sum = 0;
  set<set<int>> found;

  for (int d = 2; d < area-1; d++) {
    for (int abe = 2; abe < area-1; abe++) {
      long long bd = area - d;
      long long numer = bd * bd * abe;
      long long denom = area*area - d*abe;
      if (numer%denom == 0) {
	int bdf = numer / denom;
	int adfe = abe - bdf;
	int bfc = bd - bdf;
	int fec = area - bdf - adfe - bfc;
	set<int> areas = {bdf, adfe, bfc, fec};
	if (found.find(areas)==found.end()) {
	  sum += area;
	  found.insert(areas);
	}
      }
    }
  }
  return sum;
}

int main() {
  long long ans = 0;
  for (int i = 1; i <= 10000; i++) {
    if (i%1000 == 0) {
      cout << "At " << i << endl;
    }
    ans += findIntegralAreas(i);
  }
  cout << ans << endl;
}
