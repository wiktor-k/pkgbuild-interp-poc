# Maintainer: Orhun Parmaksız <orhun@archlinux.org>
# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: Julien Nicoulaud <julien.nicoulaud@gmail.com>

pkgname=zellij
pkgver=0.45.1
pkgrel=1
pkgdesc="A terminal multiplexer"
arch=('x86_64' 'i686' 'armv6h' 'armv7h')
url="https://zellij.dev"
license=('MIT')
depends=(curl
         glibc # libc.so libm.so
         libgcc
         zlib)
makedepends=(cargo
             mandown)
options=(!lto)
source=("https://github.com/zellij-org/${pkgname}/archive/refs/tags/v${pkgver}/${pkgname}-${pkgver}.tar.gz")
sha512sums=('1fa550fe243d8b18a87ff1983631a9a76eb9dde9d1d2c59f0aa6e1bb8b1a677168c469e143de7cc4fcc9ff3a0799c1dc566553ed74bcadab345282b81c105fc8')

_srcenv() {
  cd "$pkgname-$pkgver"
  export CARGO_HOME="$srcdir"
  export CARGO_PROFILE_RELEASE_DEBUG=2
  export CARGO_PROFILE_RELEASE_STRIP=false
  export CARGO_PROFILE_RELEASE_LTO=thin
  export CARGO_PROFILE_RELEASE_CODEGEN_UNITS=1
  export CARGO_PROFILE_RELEASE_OPT_LEVEL=3
  CFLAGS+=' -ffat-lto-objects'
}

prepare() {
  _srcenv
  cargo fetch --locked --target host-tuple
}

build() {
  _srcenv
  cargo build --release --frozen
  ./target/release/zellij setup --generate-completion bash > target/zellij.bash
  ./target/release/zellij setup --generate-completion elvish > target/zellij.elv
  ./target/release/zellij setup --generate-completion fish > target/zellij.fish
  ./target/release/zellij setup --generate-completion zsh > target/zellij.zsh
  # mandown docs/MANPAGE.md > assets/zellij.1
}

package() {
  depends+=(libcurl.so
            libgcc_s.so
            libz.so)
  cd "$pkgname-$pkgver"
  install -Dm755 target/release/zellij -t "${pkgdir}/usr/bin"
  install -Dm644 GOVERNANCE.md README.md -t "${pkgdir}/usr/share/doc/zellij"
  install -Dm644 LICENSE.md -t "${pkgdir}/usr/share/licenses/zellij"
  install -Dm644 target/zellij.bash "${pkgdir}/usr/share/bash-completion/completions/zellij"
  install -Dm644 target/zellij.elv "${pkgdir}/usr/share/elvish/lib/zellij.elv"
  install -Dm644 target/zellij.fish "${pkgdir}/usr/share/fish/vendor_completions.d/zellij.fish"
  install -Dm644 target/zellij.zsh "${pkgdir}/usr/share/zsh/site-functions/_zellij"
  # install -Dm644 assets/zellij.1 "${pkgdir}/usr/share/man/man1/zellij.1"
  install -Dm644 assets/zellij.desktop "${pkgdir}/usr/share/applications/zellij.desktop"
  install -Dm644 assets/logo.png "${pkgdir}/usr/share/pixmaps/zellij.png"
}

# vim: ts=2 sw=2 et:
