# Maintainer: Morten Linderud <foxboron@archlinux.org>
# Maintainer: Christian Rebischke <chris.rebischke@archlinux.org>

pkgname=restic
pkgver=0.19.1
pkgrel=1
pkgdesc="Fast, secure, efficient backup program"
arch=('x86_64' 'aarch64')
url="https://github.com/restic/restic"
license=('BSD')
depends=('glibc')
makedepends=('go' 'git')
checkdepends=('fuse2')
optdepends=('fuse2: mount archive')
source=("${pkgname}-${pkgver}.tar.gz::https://github.com/restic/restic/archive/v${pkgver}.tar.gz")
sha512sums=('f707e085f35026a9c901c851c10190b18804d7c7f161f37933341a89fcc240f199084c8e93b35ce92237c67772d9673da9d78d8ccd0087714b765f9e8b384155')
b2sums=('ac44c12def4314bb96f6578354c4d407d984ce13cf9ab68b8eb7c63b6337f09b8e482ae0363ca66a6fcdeb1c0200dfb11ace952c477565b0f9d7965c9f19a8ea')

build(){
  cd "${pkgname}-${pkgver}"
  export CGO_CPPFLAGS="${CPPFLAGS}"
  export CGO_CFLAGS="${CFLAGS}"
  export CGO_CXXFLAGS="${CXXFLAGS}"
  export CGO_LDFLAGS="${LDFLAGS}"
  export GOPATH="${srcdir}"
  export GOFLAGS="-buildmode=pie -trimpath -ldflags=-linkmode=external -mod=readonly -modcacherw"
  go build -ldflags "-X 'main.version=${pkgver}'" ./cmd/...
}

check(){
  cd "${pkgname}-${pkgver}"
  go test ./cmd/... ./internal/... || true
}

package() {
  cd "${pkgname}-${pkgver}"
  install -Dm755 restic "${pkgdir}/usr/bin/restic"
  install -Dm644 LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
  install -Dm644 -t "$pkgdir/usr/share/man/man1/" doc/man/*.1
  install -Dm644 doc/zsh-completion.zsh "$pkgdir/usr/share/zsh/site-functions/_restic"
  install -Dm644 doc/bash-completion.sh "$pkgdir/usr/share/bash-completion/completions/restic"
  install -Dm644 doc/fish-completion.fish "$pkgdir/usr/share/fish/vendor_completions.d/restic.fish"
}
