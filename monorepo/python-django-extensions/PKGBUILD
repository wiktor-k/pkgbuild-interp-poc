# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Jeremy Audet <ichimonji10@gmail.com>
# Contributor: Dominik Kozaczko <dominik@kozaczko.info>
# Contributor: Alper Kanat <alperkanat@raptiye.org>
# Contributor: Apkawa <apkawa@gmail.com>
# Contributor: Schnouki <thomas.jost@gmail.com>

pkgname=python-django-extensions
pkgver=4.1
pkgrel=2
pkgdesc='A collection of custom extensions for the Django Framework'
arch=(any)
url=https://github.com/django-extensions/django-extensions
license=(MIT)
depends=(
  python-django
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
optdepends=(
  'graphviz: graph_models extension'
  'ipython: shell_plus extension'
  'python-werkzeug: runserver_plus extension'
)
_tag=ad015513e5d8e8ee6ad748fca5f32ccf3fb3d5c5
source=(git+https://github.com/django-extensions/django-extensions.git#tag=${_tag})
sha256sums=('ca727cd17d56a70afe8e4e62654069e6ad325922b02c75538796e38ab97a7f75')

pkgver() {
  cd django-extensions
  git describe --tags
}

build() {
  cd django-extensions
  python -m build --wheel --no-isolation
}

package() {
  python -m installer --destdir="${pkgdir}" django-extensions/dist/*.whl
  install -Dm 644 django-extensions/LICENSE -t "${pkgdir}"/usr/share/licenses/python-django-extensions/
}

# vim: ts=2 sw=2 et:
