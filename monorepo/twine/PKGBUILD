# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: Chris Warrick <aur@chriswarrick.com>

pkgname=twine
pkgver=7.0.0
pkgrel=1
pkgdesc='Collection of utilities for interacting with PyPI'
arch=(any)
url='https://github.com/pypa/twine/'
license=(Apache-2.0)
depends=(
  python
  python-packaging
  python-readme-renderer
  python-requests
  python-requests-toolbelt
  python-urllib3
  python-importlib-metadata
  python-keyring
  python-rfc3986
  python-rich
  python-id
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-setuptools-scm
  python-wheel
)
checkdepends=(
  git
  python-munch
  python-jaraco.envs
  python-portend
  python-pretend
  python-pytest-services
  python-pytest-socket
)
provides=(python-twine)
conflicts=(python-twine)
replaces=(python-twine)
source=("$pkgname::git+https://github.com/pypa/twine.git#tag=$pkgver")
sha512sums=('644c900f1cc772b941196c0e42d7d44de7b6495e0381a7f24a98b85bdb622d48006f666b962a33defcf482c8d4c5848e8c72b122506a9ed737f278faae1f2699')
b2sums=('c6405482e9d9ba55e3b55604a5c8d69951d8c6d6131e95fd32c2494cebfbcf30de2ca25dd46e5557cd7f2b27681ad19a0020ed2b556d1c52aa886c24ef0c2461')

build() {
  cd "$pkgname"

  SETUPTOOLS_SCM_PRETEND_VERSION=$pkgver python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  # Ignored tests are copied from upstream tox.ini:
  # https://github.com/pypa/twine/blob/4.0.2/tox.ini#L16
  python -m pytest --ignore-glob '*integration*.py'
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # create symlink
  ln -s twine "$pkgdir"/usr/bin/twine3
}
