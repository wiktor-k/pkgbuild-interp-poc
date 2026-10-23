# Maintainer: Balló György <ballogyor+arch at gmail dot com>
# Maintainer: Bruno Pagani <archange@archlinux.org
# Contributor: Doug Newgard <scimmia at archlinux dot org>
# Contributor: XavierCLL <xavier.corredor.llano (a) gmail.com>
# Contributor: SaultDon <sault.don gmail>
# Contributor: Lantald < lantald at gmx.com >
# Contributor: Thomas Dziedzic < gostrc at gmail >
# Contributor: dibblethewrecker dibblethewrecker.at.jiwe.dot.org
# Contributor: Gerardo Exequiel Pozzi <vmlinuz386@yahoo.com.ar>
# Contributor: Eric Forgeot < http://esclinux.tk >

pkgname=qgis
pkgver=4.2.3
pkgrel=1
pkgdesc='Geographic Information System (GIS) that supports vector, raster & database formats'
arch=(x86_64)
url='https://qgis.org/'
license=(GPL-2.0-or-later)
depends=(
  abseil-cpp
  draco
  exiv2
  expat
  libgcc
  libstdc++
  gdal
  geos
  glibc
  gsl
  hdf5
  hicolor-icon-theme
  libspatialite
  libxml2
  libzip
  netcdf
  ocl-icd
  pdal
  postgresql-libs
  proj
  protobuf
  python
  python-gdal
  python-jinja
  python-numpy
  python-owslib
  python-packaging
  python-psycopg2
  python-pyqt6
  python-qscintilla-qt6
  python-yaml
  qca-qt6
  qscintilla-qt6
  qt6-3d
  qt6-5compat
  qt6-base
  qt6-declarative
  qt6-imageformats
  qt6-multimedia
  qt6-positioning
  qt6-serialport
  qt6-svg
  qt6-tools
  qt6-webengine
  qtkeychain-qt6
  qwt
  sqlite
  zlib
  zstd
)
makedepends=(
  cmake
  fcgi
  ninja
  opencl-clhpp
  pyqt-builder
  sip
)
optdepends=(
  'fcgi: Map server'
  'gpsbabel: GPS Tools plugin'
)
source=("https://qgis.org/downloads/$pkgname-$pkgver.tar.bz2"
        https://github.com/qgis/QGIS/commit/47ab0643e3c158476dd01dee13e7dff8eab59274.patch)
sha256sums=('b044889d895c07d9355512cba76973822830153f6f15feb0d5e00077d3d76e59'
            'ace68c93daf2616913bd34e7a9c17f4969d8dfd97ab94b1280f3246e70226ec8')

prepare() {
  cd $pkgname-$pkgver
# Fix build with Qt 6.12
  patch -p1 -i ../47ab0643e3c158476dd01dee13e7dff8eab59274.patch
}

build() {
  cmake -S $pkgname-$pkgver -B build -G Ninja \
    -DCMAKE_INSTALL_PREFIX='/usr' \
    -DWITH_3D=TRUE \
    -DWITH_QUICK=TRUE \
    -DQGIS_QML_SUBDIR=lib/qt/qml \
    -DWITH_SERVER=TRUE \
    -DWITH_CUSTOM_WIDGETS=TRUE \
    -DBINDINGS_GLOBAL_INSTALL=TRUE \
    -DQGIS_MANUAL_SUBDIR=share/man \
    -DWITH_PDAL=TRUE \
    -DWITH_INTERNAL_SPATIALINDEX=TRUE # https://github.com/libspatialindex/libspatialindex/issues/276
    # https://github.com/qgis/QGIS/issues/48374
    #-DWITH_INTERNAL_LAZPERF=FALSE \
    # https://github.com/qgis/QGIS/issues/35440
    #-DWITH_PY_COMPILE=TRUE \

  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build
  install -Dm644 $pkgname-$pkgver/rpm/sources/qgis-mime.xml "$pkgdir/usr/share/mime/packages/qgis.xml"
}
