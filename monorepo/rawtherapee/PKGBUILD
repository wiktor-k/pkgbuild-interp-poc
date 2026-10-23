# Maintainer: Christian Hesse <mail@eworm.de>
# Contributor: Lukas Jirkovsky <l.jirkovsky AT gmail.com>
# Contributor: Bogdan Szczurek <thebodzio(at)gmail.com>
# Contributor: Vaclav Kramar <vaclav.kramar@tiscali.cz>
# Contributor: Archie <mymaud@gmail.com>

pkgname=rawtherapee
epoch=1
pkgver=5.13
pkgrel=1
pkgdesc='A powerful cross-platform raw image processing program'
arch=('x86_64')
url='https://www.rawtherapee.com/'
license=('GPL-3.0-or-later')
depends=('atkmm' 'libatkmm-1.6.so'
         'cairo' 'libcairo.so'
         'cairomm' 'libcairomm-1.0.so'
         'desktop-file-utils'
         'exiv2' 'libexiv2.so'
         'expat' 'libexpat.so'
         'fftw' 'libfftw3f.so' 'libfftw3f_omp.so'
         'glib2' 'libglib-2.0.so' 'libgobject-2.0.so'
         'glibc'
         'glibmm' 'libgiomm-2.4.so' 'libglibmm-2.4.so'
         'gtk3' 'libgdk-3.so' 'libgtk-3.so'
         'gtkmm3' 'libgdkmm-3.0.so' 'libgtkmm-3.0.so'
         'hicolor-icon-theme'
         'lcms2' 'liblcms2.so'
         'lensfun' #'liblensfun.so'
         'libcanberra' 'libcanberra-gtk3.so' 'libcanberra.so'
         'libgcc' 'libgcc_s.so'
         'libgomp' 'libgomp.so'
         'libiptcdata'
         'libjpeg-turbo' 'libjpeg.so'
         'libjxl' 'libjxl.so' 'libjxl_threads.so'
         'libpng' 'libpng16.so'
         'librsvg' 'librsvg-2.so'
         'libsigc++' 'libsigc-2.0.so'
         'libstdc++' 'libstdc++.so'
         'libtiff' 'libtiff.so'
         'pangomm' 'libpangomm-1.4.so'
         'zlib' 'libz.so')
makedepends=('cmake' 'git')
source=("git+https://github.com/Beep6581/RawTherapee.git#tag=${pkgver}")
sha256sums=('52edf54be647856d1237083c998025c187fddfcb21bcc963bd0d294e72945f80')

build() {
  cmake -B build -S RawTherapee \
    -DCMAKE_INSTALL_PREFIX=/usr \
    -DCMAKE_BUILD_TYPE=None \
    -DCACHE_NAME_SUFFIX="" \
    -Wno-dev
  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}

# vim:set ts=2 sw=2 et:
