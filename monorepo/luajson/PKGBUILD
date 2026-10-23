# Maintainer: Jan Alexander Steffens (heftig) <heftig@archlinux.org>
# Contributor: Luca P <meti at lplab.net>

pkgbase=luajson
pkgname=(
  lua52-luajson
)
pkgver=1.3.5
pkgrel=1
pkgdesc="JSON parser/encoder for Lua"
url="https://github.com/harningt/luajson"
arch=(any)
license=(MIT)
makedepends=(
  git
  lua52-lpeg
)
source=("git+$url#tag=$pkgver")
b2sums=('56870424eb2cb1aa79068625bb120cbb13f80446299f97e65fb13fd3a251920335ae17b756d908cce03b3cc193ceb613d0302f4a3a8f677fe44a7858b7f71a54')

package_lua52-luajson() {
  pkgdesc+=" 5.2"
  depends=(
    lua52-lpeg
  )

  make -C $pkgbase PREFIX=/usr DESTDIR="$pkgdir" \
    INSTALL_LMOD='$(INSTALL_TOP)/share/lua/5.2' \
    INSTALL_CMOD='$(INSTALL_TOP)/lib/lua/5.2' install

  install -Dm644 $pkgbase/LICENSE -t "$pkgdir/usr/share/licenses/$pkgname"
}

# vim:set sw=2 sts=-1 et:
