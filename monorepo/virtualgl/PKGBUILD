# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
# Contributor: FigoDaPaura <ffigoDaPaura>
# Contributor: Andres Jimenez

pkgname=virtualgl
pkgver=3.1.5
pkgrel=1
pkgdesc="Redirects 3D commands from an OpenGL application onto a server-side 3D graphics card"
arch=('x86_64')
url="https://virtualgl.org"
license=('LGPL')
depends=('libxv' 'glu' 'libjpeg-turbo' 'xcb-util-keysyms' 'opencl-driver' 'ocl-icd' 'xorg-xhost' 'inetutils')
makedepends=('cmake' 'mesa' 'libxtst' 'opencl-headers' 'ninja')
source=("https://github.com/VirtualGL/virtualgl/releases/download/${pkgver}/VirtualGL-${pkgver}.tar.gz"{,.sig}
        "$pkgname.sysusers")
sha512sums=('c9b138b1c7f979548f07c66025cbde7592d21ea23da2415ed6a715708d44fc290c052fe484cd521b3a277e7c6d0fcadecff0cea3ec7716b6f65cff6c3c1ba5d0'
            'SKIP'
            'fa97a6c40a216db1632563187bbc5452025cfb8f93341b84e8d71894113f4236f4c64cb38bdd0c4ac55eececec2aad6f357b8716af08b181d3f1352da34fe799')
validpgpkeys=('AE1A7BA4EFFF9A9987E1474C4BACCAB36E7FE9A1') # The VirtualGL Project (Signing key for official binaries) <information@VirtualGL.org>

prepare() {
  cd "VirtualGL-$pkgver"

  # FS#48403
  export LDFLAGS="${LDFLAGS/--as-needed,}"
}

build() {
  cd "VirtualGL-$pkgver"

  cmake \
      -GNinja \
      -Bbuild \
      -DCMAKE_INSTALL_PREFIX=/usr \
      -DTJPEG_INCLUDE_DIR=/usr/include \
      -DTJPEG_LIBRARY=/usr/lib/libturbojpeg.so \
      -DCMAKE_INSTALL_LIBDIR=/usr/lib/virtualgl \
      -DCMAKE_INSTALL_BINDIR=/usr/bin \
      -DCMAKE_INSTALL_DOCDIR=/usr/share/doc/${pkgname} \
      -DVGL_FAKEXCB=1
  ninja -C build
}

package() {
  cd "VirtualGL-$pkgver"
  DESTDIR="$pkgdir" ninja -C build install

  mv "$pkgdir"/usr/bin/glxinfo "$pkgdir"/usr/bin/vglxinfo
  mv "$pkgdir"/usr/bin/eglinfo "$pkgdir"/usr/bin/veglinfo
  install -Dm0644 "${srcdir}/${pkgname}.sysusers" "${pkgdir}/usr/lib/sysusers.d/${pkgname}.conf"
}
