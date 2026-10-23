# Maintainer: Filipe Laíns (FFY00) <lains@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=arduino-cli
pkgver=1.5.1
pkgrel=1
arch=('x86_64')
pkgdesc='Arduino command line interface'
url='https://github.com/arduino/arduino-cli'
license=('GPL-3.0-only')
groups=('arduino')
depends=('glibc')
makedepends=(
  'git'
  'go'
)
source=("git+$url.git#tag=v$pkgver")
b2sums=('9e57cabb66b018f3771f729e3c0004ff06cd8f8540d5fec3cc295679dcad2aeeae165f27c07222be872dfbe29e7de68209def7e97863cbb19ad2c7b7a7e291a3')

prepare() {
  cd $pkgname
  GOFLAGS="-mod=readonly" go mod vendor -v
}

build(){
  cd $pkgname
  export CGO_LDFLAGS="$LDFLAGS"
  export CGO_CFLAGS="$CFLAGS"
  export CGO_CPPFLAGS="$CPPFLAGS"
  export CGO_CXXFLAGS="$CXXFLAGS"
  export GOFLAGS="-buildmode=pie -mod=vendor -modcacherw -buildvcs=false"
  export GOPATH="$srcdir"

  local ld_flags=" \
    -X github.com/arduino/arduino-cli/internal/version.versionString=$pkgver \
    -X github.com/arduino/arduino-cli/internal/version.commit=$(git rev-parse HEAD) \
    -compressdwarf=false \
    -linkmode=external \
  "
  go build -v -ldflags "$ld_flags" -o build/
}

check() {
  cd $pkgname
  local unit_tests=$(
    go list ./... \
      | grep -v github.com/arduino/arduino-cli/internal/arduino/monitor \
      | grep -v github.com/arduino/arduino-cli/internal/integrationtest \
      | grep -v github.com/arduino/arduino-cli/internal/version \
  )
  # shellcheck disable=SC2086
  go test -v $unit_tests
}

package() {
  cd $pkgname
  install -vDm755 -t "$pkgdir/usr/bin" build/$pkgname
}
