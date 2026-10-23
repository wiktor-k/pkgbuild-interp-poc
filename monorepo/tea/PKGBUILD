# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Maintainer: Robin Candau <antiz@archlinux.org>
# Contributor: John K. Luebs <jkl at johnluebs dot tld>

pkgname=tea
pkgver=0.16.0
pkgrel=1
pkgdesc='A command line tool to interact with Gitea servers'
arch=(x86_64)
url='https://gitea.com/gitea/tea'
license=(MIT)
depends=(glibc)
makedepends=(git go)
replaces=(gitea-tea)
options=(!lto)
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('669ddbedda1193e0802b9c4b685f90afd50d02495733060b67e57dab033bbf858ab5c0e925ed64a158daac0003501b91d0857d449a77b32b93127cc50487fba3')
b2sums=('7a8cfe7ea8ef274158ba5631f201b580e0b0b60a109dbec5ecdaa20587223952adbb4fee83075dba0dbfeaf80463dfb33317e45e617bbcc4dab3637ad701c0fc')

prepare() {
  cd "$pkgname"

  # create directory for build output
  mkdir build

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
  local TAGS=""

  go build -v \
    -buildmode=pie \
    -mod=readonly \
    -modcacherw \
    -ldflags "-compressdwarf=false \
    -linkmode external \
    -extldflags '${LDFLAGS}' \
    -X gitea.dev/tea/modules/version.Version=${pkgver} \
    -X gitea.dev/tea/modules/version.Tags=${TAGS}" \
    -o build \
    .

  # generate shell completions
  for shell in bash fish zsh; do
    build/tea completion "$shell" > build/"${shell}-completion"
  done

  # generate man page
  build/tea man --out build/tea.1
}

package() {
  cd "$pkgname"

  # binary
  install -vDm755 -t "$pkgdir/usr/bin" build/tea

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE

  # shell completions
  install -vDm644 build/bash-completion \
    "$pkgdir/usr/share/bash-completion/completions/tea"
  install -vDm644 build/fish-completion \
    "$pkgdir/usr/share/fish/vendor_completions.d/tea.fish"
  install -vDm644 build/zsh-completion \
    "$pkgdir/usr/share/zsh/site-functions/_tea"

  # man page
  install -vDm644 build/tea.1 \
    "$pkgdir/usr/share/man/man1/tea.1"
}
