# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Contributor : Martin Wimpress <code@flexion.org>

pkgname=gufw
pkgver=26.04
pkgrel=2
pkgdesc='Uncomplicated way to manage your Linux firewall'
url='https://costales.github.io/projects/gufw/'
arch=('any')
license=('GPL3')
depends=('gtk3' 'polkit' 'python' 'python-gobject' 'ufw' 'webkit2gtk-4.1'
         'desktop-file-utils' 'gtk-update-icon-cache')
makedepends=('git' 'intltool' 'python-distutils-extra' 'python-setuptools')
source=(git+https://github.com/costales/gufw.git#tag=${pkgver})
sha512sums=('d68d797265f276bbcd01929733b3248f82815dee0f7feb5ef4a58d25adc6748f01e0f484042ca28be73e1f17a80aefb46896a83df985d2c96c65265306bfa9c6')
b2sums=('48d88dca594bf386aba21bfbce6c16b4fb04a31fae738f2cce4d323671a44f1adaca05bf2dcbfd90f17c265871828cea9a8dcefa5086458a21c7547ed171d266')

prepare () {
  cd ${pkgname}

  local PYTHONVERSION="$(python -c 'import sys; print("{}.{}".format(sys.version_info.major, sys.version_info.minor))')"
  sed 's|/usr/share/gufw/gufw/gufw.py $1|/usr/lib/python'"${PYTHONVERSION}"'/site-packages/gufw/gufw.py "$@"|' -i bin/gufw-pkexec
}

build() {
  cd ${pkgname}
  python setup.py build
}

package() {
  cd ${pkgname}
  # required so desktop, icon and translation files get installed
  export XDG_RUNTIME_DIR="$(mktemp -d --tmpdir gufw-build-XXXXXXXXX)"
  python setup.py install --prefix=/usr --root="${pkgdir}" -O1
}

# vim: ts=2 sw=2 et:
