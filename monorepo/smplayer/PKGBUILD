# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: Ionut Biru <ibiru@archlinux.org>
# Contributor: Allan McRae <allan@archlinux.org>
# Contributor: Alessio 'mOLOk' Bolognino <themolok@gmail.com>
# Contributor: shamrok <szamrok@gmail.com>

pkgname=smplayer
pkgver=25.6.0
pkgrel=2
pkgdesc='Media player with built-in codecs that can play virtually all video and audio formats'
url='https://www.smplayer.info/'
arch=(x86_64)
license=(GPL-2.0-or-later)
depends=(
  glibc
  hicolor-icon-theme
  libgcc
  libstdc++
  libx11
  mpv
  qt5-base
  qt5-declarative
  zlib
)
makedepends=(
  qt5-tools
)
optdepends=(
  'smplayer-themes: icon themes collection'
  'smplayer-skins: skin themes collection'
# 'smtube: browse and play youtube videos'
  'youtube-dl: youtube videos and streaming'
  'yt-dlp: youtube videos and streaming'
  'mplayer: alternative multimedia engine'
)
source=(
  https://github.com/smplayer-dev/${pkgname}/releases/download/v$pkgver/${pkgname}-${pkgver}.tar.bz2
)
sha512sums=('835f0faad14fe94b019f38d5e5ef34ff490274c94818677ee8e44cf6a74024eadd8173de98bc2e6be3b5df97843908bfaf500f1a007aed9bb34c22a91e75d7db')
b2sums=('ed250381902aa8c905da8502af7c4d4e977b827ccd2bece385fe435500ab2cfd175dce11a8603a638ad25123fd8cc2cd19eee28d371a7abd3a248bbd557d2c1c')

prepare() {
  cd ${pkgname}-${pkgver}
  # make build reproducible, we compress ourselves
  sed '/gzip -9/d' -i Makefile
}

build() {
  cd ${pkgname}-${pkgver}
  export CXXFLAGS="${CXXFLAGS} ${CPPFLAGS}"
  make \
    PREFIX=/usr \
    DOC_PATH="\\\"/usr/share/doc/smplayer\\\"" \
    QMAKE_OPTS=DEFINES+=NO_DEBUG_ON_CONSOLE \
    CFLAGS_EXTRA="${CFLAGS} ${CPPFLAGS} ${LDFLAGS}" \
    all
}

package() {
  cd ${pkgname}-${pkgver}
  make \
    PREFIX=/usr \
    DOC_PATH=/usr/share/doc/smplayer \
    DESTDIR="${pkgdir}" \
    install
}

# vim: ts=2 sw=2 et:
