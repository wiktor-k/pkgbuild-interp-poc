# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Arthur Zamarin <arthurzam@gmail.com>
# Contributor: andrew.p <pan.pav.7c5@gmail.com>

pkgname=lightdm-gtk-greeter-settings
pkgver=1.2.3
pkgrel=5
pkgdesc='Settings editor for the LightDM GTK+ Greeter'
arch=(any)
url=https://github.com/xubuntu/lightdm-gtk-greeter-settings
license=(GPL-3.0-only)
depends=(
  bash
  gdk-pixbuf2
  glib2
  gtk-update-icon-cache
  gtk3
  hicolor-icon-theme
  lightdm-gtk-greeter
  pango
  python
  python-gobject
)
makedepends=(
  git
  python-build
  python-distutils-extra
  python-installer
  python-setuptools
  python-wheel
)
source=(git+https://github.com/Xubuntu/lightdm-gtk-greeter-settings.git#tag=lightdm-gtk-greeter-settings-${pkgver})
b2sums=('e9fe51dcb33c1f37188b7e50405f4b16284ea46104a543f8e8d2c3fec8741965cece8e9c9ece22315845348e31391449f33a19332c8a043d2fb433f2fc4f8667')

prepare() {
  cd lightdm-gtk-greeter-settings
  # Add Xfce pluggable support
  # Manual implementation of the --xfce-integration flag logic which is not compatible with PEP517
  echo "X-XfcePluggable=true" >> lightdm-gtk-greeter-settings.desktop.in
  # Fix __data_directory__ value in installation_config.py when using PEP517
  sed -i "s|\(target_pkgdata = \)target_data|\1'/usr/'|" setup.py
}

build() {
  cd lightdm-gtk-greeter-settings
  python -m build --wheel --no-isolation
}

package() {
  cd lightdm-gtk-greeter-settings
  python -m installer --destdir="${pkgdir}" dist/*.whl
}
