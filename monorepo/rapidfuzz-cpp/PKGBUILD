# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Pekka Ristola <pekkarr [at] protonmail [dot] com>

pkgname=rapidfuzz-cpp
pkgver=3.3.4
pkgrel=1
pkgdesc='Rapid fuzzy string matching in C++ using the Levenshtein Distance'
arch=(any)
url='https://github.com/rapidfuzz/rapidfuzz-cpp'
license=(MIT)
depends=(cmake)
makedepends=(git catch2)
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('6754afed21a36fb7e4bc7534aad69178d20ed8dc4132c112d39f5a9378aaadc93d7918d8730e9bfcf30c699fc06a2d71c2e914f13ab40ed3ac763a2459821ae9')
b2sums=('a0c87d2dd5daad15bb493330991d16b879506fc31148ba5c54081acbade722e1ebce96bdc3020091b2a07af575dc40510e30386dae340d2051dcf59d99cc5252')

build() {
  cmake \
    -B build \
    -S "$pkgname" \
    -D RAPIDFUZZ_BUILD_TESTING=ON \
    -D RAPIDFUZZ_ENABLE_LINTERS=ON \
    -D CMAKE_BUILD_TYPE=None \
    -D CMAKE_INSTALL_PREFIX=/usr \
    -W no-dev

  cmake --build build
}

check() {
  ctest --test-dir build --output-on-failure
}

package() {
  DESTDIR="$pkgdir" cmake --install build

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" "$pkgname/LICENSE"
}
