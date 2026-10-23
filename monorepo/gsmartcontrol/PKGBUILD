# Maintainer: Jelle van der Waa <jelle@vdwaa.nl>
# Maintainer: Orhun Parmaksız <orhun@archlinux.org>

pkgname=gsmartcontrol
pkgver=2.0.2
pkgrel=1
pkgdesc="A graphical user interface for the smartctl hard disk drive health inspection tool."
arch=('x86_64')
url="https://gsmartcontrol.sourceforge.io/home/"
license=('GPL')
depends=('smartmontools' 'gtkmm3' 'xorg-xmessage' 'hicolor-icon-theme')
optdepends=('polkit: to run gsmartcontrol directly from menu'
  'xterm: to update the drive database')
makedepends=('cmake')
source=("$pkgname-$pkgver.tar.gz::https://github.com/ashaduri/gsmartcontrol/archive/v$pkgver.tar.gz"
  org.archlinux.pkexec.gsmartcontrol.policy
  gsmartcontrol_polkit)
sha256sums=('7cebd83fd34883d51e143389aa88f8173ea7b67c760b12b7de847f3c3c8cee34'
            '98ae73cbe39be31f885d8c4d09cbfac7b454534ea1ba7cb5a8d23365dc22212a'
            '66f952196db24d6e3ee454b03481d52e0281c3bbf8be180600bf5ba7b39ec0ea')

prepare() {
  cd ${pkgname}-${pkgver}

  # Modify desktop file
  sed -i "s|^Exec=.*|Exec="@prefix@/bin/gsmartcontrol_polkit"|" \
    data/gsmartcontrol.in.desktop
}

build() {
  cd ${pkgname}-${pkgver}
  mkdir build
  cd build
  cmake .. -DCMAKE_INSTALL_PREFIX=/usr -DCMAKE_INSTALL_SBINDIR=/usr/bin -DCMAKE_BUILD_TYPE=Release
  make
}

package() {
  cd ${pkgname}-${pkgver}

  mkdir -p "${pkgdir}/usr/bin"

  # Install launcher script
  install -m755 "${srcdir}"/gsmartcontrol_polkit \
    "${pkgdir}/usr/bin/gsmartcontrol_polkit"

  # Install policy file
  install -Dm644 "${srcdir}"/org.archlinux.pkexec.gsmartcontrol.policy \
    "${pkgdir}/usr/share/polkit-1/actions/org.archlinux.pkexec.gsmartcontrol.policy"

  cd build
  make DESTDIR="${pkgdir}" install

}
