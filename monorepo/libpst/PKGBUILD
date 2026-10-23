# Maintainer: Jan Alexander Steffens (heftig) <heftig@archlinux.org>
# Contributor: Jan de Groot <jgc@archlinux.org>

pkgbase=libpst
pkgname=(
  libpst
  libpst-docs
)
pkgver=0.6.76
pkgrel=11
pkgdesc="Outlook .pst file converter"
url="https://www.five-ten-sg.com/libpst/"
arch=(x86_64)
license=(GPL-2.0-or-later)
depends=(
  glib2
  glibc
  libgcc
  libgsf
  libstdc++
  zlib
)
makedepends=(
  autoconf-archive
  docbook-xsl
  doxygen
  git
  graphviz
  xmlto
)
source=("git+https://github.com/pst-format/libpst#tag=libpst-$pkgver")
b2sums=('0644efe8a4b270b891aaf4df72f19d0d47fb074ee7f3c5654182b72923ba2545b02132ba2f0cf65b2776d5fbaa906263f52d6f4177ad7964c9d33d8d6c19de37')

prepare() {
  cd $pkgbase

  # Build fixes
  git cherry-pick -n df169eaeaa3a0aad7244c081497a01ceee15da72 \
                     81252ea84b5a6a1e0d3cbcfaf4505729a1057018 \
                     db5933678ac44f9153d0c3dd20a6cb91a6887096 \
                     52adecb804f4d8fa1697692e177e972eb891fcb2

  # Fix _FORTIFY_SOURCE warnings
  git cherry-pick -n b34d3d8eb679e3598dd9970bb35e6ab60fdd9805

  # Fix build
  git cherry-pick -n cc600ee98c4ed23b8ab0bc2cf6b6c6e9cb587e89

  autoreconf -fvi
}

build() {
  local configure_options=(
    --prefix=/usr
    --sysconfdir=/etc
    --localstatedir=/var
    --mandir=/usr/share/man
    --disable-dii
    --disable-python
    --disable-static
    --enable-libpst-shared
  )

  cd $pkgbase
  ./configure "${configure_options[@]}"
  sed -i -e 's/ -shared / -Wl,-O1,--as-needed\0/g' libtool
  make -C xml all
  make
}

package_libpst() {
  provides=(libpst.so)

  make -C $pkgbase DESTDIR="$pkgdir" install

  mkdir -p doc/usr/share
  mv {"$pkgdir",doc}/usr/share/doc
}

package_libpst-docs() {
  pkgdesc+=" (documentation)"
  depends=()

  mv doc/* "$pkgdir"
}

# vim:set sw=2 sts=-1 et:
