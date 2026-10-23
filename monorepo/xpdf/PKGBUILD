# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Hyacinthe Cartiaux <hyacinthe.cartiaux@free.fr>
# Contributor: Gaetan Bisson <bisson@archlinux.org>
# Contributor: tobias <tobias@archlinux.org>
# Contributor: Sarah Hay <sarahhay@mb.sympatico.ca>
# Contributor: Mark Pustjens <pustjens@dds.nl>

pkgname=xpdf
pkgver=4.06
pkgrel=2
pkgdesc='Viewer for Portable Document Format (PDF) files'
arch=('x86_64')
url='https://www.xpdfreader.com/'
license=('GPL-2.0-only OR GPL-3.0-only')
depends=(
  'fontconfig'
  'freetype2'
  'glibc'
  'gsfonts'
  'libcups'
  'libgcc'
  'libjpeg-turbo'
  'libpaper'
  'libpng'
  'libstdc++'
  'libtiff'
  'libxt'
  'qt6-base'
  'qt6-svg'
  'xdg-utils'
)
makedepends=(
  'cmake'
  'qt6-tools'
  'zlib'
)
optdepends=(
  'desktop-file-utils: for desktop environments'
  'poppler: tools that used to be included in xpdf'
)
backup=('etc/xpdfrc')
source=(
  "https://dl.xpdfreader.com/xpdf-${pkgver}.tar.gz"
  "https://dl.xpdfreader.com/xpdf-${pkgver}.tar.gz.sig"
  "https://src.fedoraproject.org/rpms/xpdf/raw/rawhide/f/xpdf-4.06-permissions.patch"
  "xpdf-4.00-xpdfrc.patch"
  "xpdf.desktop"
)
b2sums=('cd1d6a839bff4b13fb4461466600823f786057bcd782ba518eb788ac5dc1a6c539fd07208ff95718d48153c4d58ada6519c7f830a6c315e82987f3ed4ef9dd34'
        'SKIP'
        'cf2b5adcbd306b0fd4dfe1f60d5d93eca26a79c2a7ccf603adb16d6df4ccb2879773254b2078aa618fdd7f49e4c877eafe6f7f01e99616d6bdc60d88bb648ab6'
        '5f137a967b6631813f16e68025dd58423148faffdfe41bb3b6d4c5cdd5ed6325dca33598c2c9365be423f34a9928f12b5174e27cf47ec8cfe0c8df683a077ebc'
        'dbdb27ef76a10608aa1d265f947a8be9b20840dbb2c2c4a547883cebe084512dbabf9555b33e0e4689b610078141b501287cc926b1f2bbe14e68925c0459eb4c')
validpgpkeys=('A56006CA75CF8B13FA2F120DF4825F5397271342') # -Xpdf- <xpdf@xpdfreader.com>

prepare() {
  cd ${pkgname}-${pkgver}
  cp doc/sample-xpdfrc xpdfrc
  patch -Np1 < ../xpdf-4.06-permissions.patch
  patch -Np1 < ../xpdf-4.00-xpdfrc.patch
  sed 's:/usr/share/fonts/type1/gsfonts:/usr/share/fonts/Type1:' -i xpdf/GlobalParams.cc
}

build() {
  cmake -S ${pkgname}-${pkgver} -B build \
    -DCMAKE_BUILD_TYPE=None \
    -DCMAKE_INSTALL_PREFIX=/usr \
    -Wno-dev \
    -DMULTITHREADED=ON \
    -DOPI_SUPPORT=ON \
    -DXPDFWIDGET_PRINTING=ON \
    -DSYSTEM_XPDFRC=/etc/xpdfrc
  cmake --build build
}

package() {
  DESTDIR="${pkgdir}" cmake --install build

  cd ${pkgname}-${pkgver}
  install -vDm 644 -t "${pkgdir}/etc" xpdfrc
  install -vDm 644 -t "${pkgdir}/usr/share/applications" ../xpdf.desktop
  install -vDm 644 xpdf-qt/xpdf-icon.svg "${pkgdir}/usr/share/pixmaps/xpdf.svg"
  install -vDm 644 -t "${pkgdir}/usr/share/doc/${pkgname}" README CHANGES

  # remove stuff provided by poppler
  for tool in pdfdetach pdffonts pdfimages pdfinfo pdftoppm pdftops pdftotext pdftohtml; do
    rm -v "${pkgdir}/usr/bin/${tool}" "${pkgdir}/usr/share/man/man1/${tool}.1"
  done
}
