# Maintainer: kpcyrd <kpcyrd[at]archlinux[dot]org>
# Contributor: Francois Boulogne <devel at sciunto dot org>

pkgname=mat2
pkgver=0.15.0
pkgrel=1
pkgdesc='Metadata removal tool, supporting a wide range of commonly used file formats'
url='https://github.com/jvoisin/mat2'
arch=('any')
license=('LGPL-3.0-or-later')
depends=(
  'gdk-pixbuf2'
  'hicolor-icon-theme'
  'librsvg'
  'mailcap'
  'perl-image-exiftool'
  'poppler-glib'
  'python'
  'python-cairo'
  'python-gobject'
  'python-mutagen'
  'zenity'
)
makedepends=(
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-wheel'
)
optdepends=(
  'ffmpeg: video support'
  'bubblewrap: sandboxing support'
  'kdialog: dolphin file browser integration'
)
checkdepends=(
  'ffmpeg'
  'python-pytest'
)
source=(${pkgname}-${pkgver}.tar.gz::https://github.com/jvoisin/${pkgname}/archive/refs/tags/${pkgver}.tar.gz
        ${pkgname}-${pkgver}.tar.gz.asc::https://github.com/jvoisin/${pkgname}/releases/download/${pkgver}/${pkgname}-${pkgver}.tar.gz.asc)
sha512sums=('b15d06e62438bd0909066af8b0a13772113670eeeefb9be59123a700ac9357b44ab59f79b45340d1122dcdab86334de44ea84abb4dc1968cdd9362413c6c44b9'
            'SKIP')
b2sums=('a2aeb62e8eb12ed88886352a6aa339b1505928ec4af6dceabdbf50925978e0fe976c8080580cb5faf109efefd3257f1a5c31c6f6a0087c600304df3facf15f13'
        'SKIP')

validpgpkeys=('9FCDEE9E1A381F311EA62A7404D041E8171901CC') # Julien (jvoisin) Voisin

build() {
  cd ${pkgname}-${pkgver}
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname}-${pkgver}
  pytest -vv --deselect tests/test_libmat2.py::TestCleaning::test_all_parametred \
    --deselect tests/test_climat2.py::TestControlCharInjection::test_jpg \
    --deselect tests/test_corrupted_files.py::TestCorruptedFiles::test_wrong_tif
}

package() {
  cd ${pkgname}-${pkgver}
  python -m installer --destdir="${pkgdir}" dist/*.whl
  install -vDm 644 dolphin/mat2.desktop -t "${pkgdir}/usr/share/kio/servicemenus"
  install -vDm 644 nemo/mat2.nemo_action -t "${pkgdir}/usr/share/nemo/actions"
  install -vDm 644 data/mat2.svg -t "${pkgdir}/usr/share/icons/hicolor/scalable/apps"
  install -vDm 644 doc/*.md -t "${pkgdir}/usr/share/doc/${pkgname}"
}

# vim: ts=2 sw=2 et:
