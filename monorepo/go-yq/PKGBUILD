# Maintainer: Daniel M. Capella <polyzen@archlinux.org>
# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: Mario Finelli <mario at finel dot li>

pkgname=go-yq
pkgver=4.54.1
pkgrel=1
pkgdesc='Portable command-line YAML processor'
arch=(x86_64)
url=https://github.com/mikefarah/yq
license=(MIT)
depends=(glibc)
makedepends=(git
             go)
conflicts=(yq)
source=("$pkgname::git+$url.git#tag=v$pkgver")
sha256sums=('48d42fcbd6fa84b9d92cf65a6b3d2267cc61ec278ebdd6cbd431d98d72307b8a')

prepare() {
  cd "$pkgname"
  go mod vendor
}

build() {
  cd "$pkgname"
  export CGO_CPPFLAGS="$CPPFLAGS"
  export CGO_CFLAGS="$CFLAGS"
  export CGO_CXXFLAGS="$CXXFLAGS"
  export CGO_LDFLAGS="$LDFLAGS"
  export GOFLAGS="-buildmode=pie -trimpath -ldflags=-linkmode=external -mod=readonly -modcacherw"
  go build .
}

check() {
  cd "$pkgname"
  go test ./...
}

package() {
  cd "$pkgname"
  install -Dt "$pkgdir"/usr/bin yq
  install -Dm644 -t "$pkgdir"/usr/share/licenses/$pkgname LICENSE
  install -Dm644 -t "$pkgdir"/usr/share/doc/$pkgname README.md

  ./yq shell-completion bash | install -Dm644 /dev/stdin \
    "$pkgdir"/usr/share/bash-completion/completions/yq
  ./yq shell-completion zsh | install -Dm644 /dev/stdin \
    "$pkgdir"/usr/share/zsh/site-functions/_yq
  ./yq shell-completion fish | install -Dm644 /dev/stdin \
    "$pkgdir"/usr/share/fish/vendor_completions.d/yq.fish
}

# vim:set ts=2 sw=2 et:
