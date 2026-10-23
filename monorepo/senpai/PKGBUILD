# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgname=senpai
pkgver=0.5.0
pkgrel=2
pkgdesc='A modern terminal IRC client'
arch=(x86_64)
url='https://git.sr.ht/~delthas/senpai'
license=(ISC)
depends=(glibc)
makedepends=(git go scdoc)
optdepends=('harper: spell checking')
options=(!lto)
source=("git+$url#tag=v$pkgver")
sha512sums=('ce0332fc595a24218cddd3a824a2832d194dd7f19fa136e5a44f5e62300f240251e37fadfa4585363d2f4098d345dc1c017fc74e2ef3de1842d276f56b9b1db2')
b2sums=('34a59431c5eb9647aa844eb8137166ac7345cb49d22b585a7d2c9574c90d5838590e892be9c35b0f8ab4863276e8218a8017f7d986159ae9c9b1b26c92c96bdd')

prepare() {
  cd "$pkgname"

  # download dependencies
  export GOPATH="${srcdir}"
  go mod download
}

build() {
  cd "$pkgname"

  # set Go flags
  export CGO_CPPFLAGS="${CPPFLAGS}"
  export CGO_CFLAGS="${CFLAGS}"
  export CGO_CXXFLAGS="${CXXFLAGS}"
  export GOPATH="${srcdir}"

  go build -v \
    -buildmode=pie \
    -mod=readonly \
    -modcacherw \
    -ldflags "-compressdwarf=false \
    -linkmode external \
    -extldflags '${LDFLAGS}'" \
    ./cmd/senpai

  # documentation
  make doc/senpai.{1,5}

  # ensure reproducibility of man pages
  # i'm not 100% sure where the man sub-command gets the date from, i assume
  # it uses $TODAYS_DATE, so this should make the man page reproducible
  local _commit_date=$(git show --no-patch --format=%cd --date=format:%Y-%m-%d)
  sed \
    -i doc/*.{1,5} \
    -e "s/\"[0-9]\{4\}-[0-9]\{2\}-[0-9]\{2\}\"/\"$_commit_date\"/"
}

check() {
  cd "$pkgname"

  go test -v ./...
}

package() {
  cd "$pkgname"

  make DESTDIR="$pkgdir" PREFIX=/usr install

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
