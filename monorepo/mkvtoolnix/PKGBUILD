# Maintainer: Giovanni Scafora <giovanni@archlinux.org>
# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: xduugu <xduugu@gmx.com>

pkgbase=mkvtoolnix
pkgname=(
  mkvtoolnix-cli
  mkvtoolnix-gui
)
pkgver=102.0
pkgrel=1
pkgdesc='Set of tools to create, edit and inspect Matroska files'
arch=(x86_64)
url=https://mkvtoolnix.download/
license=(GPL-2.0-only)
depends=(
  boost-libs
  flac
  fmt
  glibc
  gmp
  libdvdread
  libgcc
  libebml
  libmatroska
  libvorbis
  libstdc++
  pugixml
  qt6-base
)
makedepends=(
  boost
  cmark
  docbook-xsl
  git
  po4a
  qt6-multimedia
  qt6-svg
  qt6-tools
  ruby-erb
  ruby-rake
  zlib
)
source=(git+https://codeberg.org/mbunkus/mkvtoolnix.git?signed#tag=release-${pkgver})
validpgpkeys=(D9199745B0545F2E8197062B0F92290A445B9007) # Moritz Bunkus <moritz@bunkus.org>
b2sums=('510ede1023c28a49f5e5615a5043c73ff527878c8634f4b655fb840ae322d50013af216f5414cfcd07131d65b9ef76b55e21df4ebd1e49a84941b25a5e1eff06')

prepare() {
  cd mkvtoolnix
  ./autogen.sh
}

build() {
  cd mkvtoolnix
  ./configure \
    --prefix=/usr \
    --disable-update-check
  rake $MAKEFLAGS
}

package_mkvtoolnix-cli() {
  depends+=(
    libogg
    libvorbis
    zlib
  )

  cd mkvtoolnix
  rake DESTDIR="${pkgdir}" install
  rm -rf "${pkgdir}"/usr/share/{applications,icons,metainfo,mime,mkvtoolnix}
  find "${pkgdir}" -name mkvtoolnix-gui* -delete
}

package_mkvtoolnix-gui() {
  depends+=(
    cmark
    hicolor-icon-theme
    mkvtoolnix-cli
    qt6-multimedia
    qt6-svg
  )

  cd mkvtoolnix
  rake DESTDIR="${pkgdir}" install
  rm -rf "${pkgdir}"/usr/share/locale
  for t in mkv{extract,info,merge,propedit}; do
    find "${pkgdir}" -name $t* -delete
  done
}
