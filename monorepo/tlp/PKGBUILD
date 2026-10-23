# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Marc Schulte <bomba@nerdstube.de>

pkgbase=tlp
pkgname=(
  tlp
  tlp-pd
  tlp-rdw
)
pkgver=1.10.2
pkgrel=1
arch=(any)
url=https://linrunner.de/en/tlp/tlp.html
license=(GPL-2.0-or-later)
makedepends=(git)
source=(git+https://github.com/linrunner/TLP.git#tag=${pkgver})
sha256sums=('614c941f8c269c0b47b269d974ec817d952956b4199c6ce878c4bcbabf509012')
install=tlp.install


package_tlp() {
  pkgdesc='Linux Advanced Power Management'
  depends=(
    bash
    hdparm
    iw
    pciutils
    perl
    rfkill
    usbutils
  )
  optdepends=(
    'bash-completion: Bash completion'
    'ethtool: Disable Wake On Lan'
    'smartmontools: Display S.M.A.R.T. data in tlp-stat'
    'tp_smapi: Older ThinkPad battery functions (before Sandy Bridge)'
  )
  conflicts=(tuned)
  backup=(etc/tlp.conf)

  export TLP_NO_INIT=1
  export TLP_SBIN=/usr/bin
  export TLP_WITH_ELOGIND=0
  export TLP_WITH_SYSTEMD=1

  make DESTDIR="${pkgdir}" -C TLP install-tlp install-man-tlp
}

package_tlp-pd() {
  pkgdesc='Linux Advanced Power Management - Power Profiles Daemon'
  depends=(
    glib2
    python
    python-gobject
    python-shtab
    tlp
  )
  provides=(power-profiles-daemon)
  conflicts=(
    power-profiles-daemon
    tuned-ppd
  )

  make DESTDIR="${pkgdir}" -C TLP install-pd install-man-pd
}

package_tlp-rdw() {
  pkgdesc='Linux Advanced Power Management - Radio Device Wizard'
  depends=(
    bash
    networkmanager
    tlp
  )

  make DESTDIR="${pkgdir}" -C TLP install-rdw install-man-rdw
}

# vim: ts=2 sw=2 et:
