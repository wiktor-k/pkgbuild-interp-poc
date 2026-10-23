# Maintainer: Lukas Fleischer <lfleischer@archlinux.org>
# Contributor: Andrey Mivrenik <gim at fastmail dot fm>
# Contributor: Tim Diels <tim@timdiels.be>

pkgname=python-dropbox
pkgver=12.2.3
pkgrel=1
pkgdesc="Python SDK for Dropbox Core APIs"
url="https://github.com/dropbox/dropbox-sdk-python"
arch=(any)
license=(MIT)
depends=(
  python
  python-requests
  python-stone
  python-urllib3
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
  python-ply
  python-pytest
  python-pytest-mock
)
source=(
  "$pkgname::git+$url#tag=v$pkgver"
  remove-version-constraints.patch
)
sha512sums=('95fb9ca398c029e88a93a8dab273e97074f3f0c203d0efb0d1b467695c0be9382b479e0d908f9aa19eeb49b86987311ad28745cba95c0705407e88f593d3bd95'
            '34538c6664424b4e5b2b73f202518db96dea67be8ac76db36dc2464d4f2e5d03dbc174d544993ebf3b093a05b7a8358155e3e0813393e138cc78d08c20759f82')
b2sums=('130bb49c6859dcfe7e26e458141629cd4642a696f48fd8028c34ce4e6a85bdaa4880de208d906d1d53fc68f7879e2dafc40e87ed03c06a79a978d861672c6626'
        '522bae5dd90c4abf9fe18b8ab0069cc1707f8759c19fd5821bdb7b4275db7c7c32513cd6d2e9932976ec99302a091cf8b347c562ed5a978b695a71edc7d21daa')

prepare() {
  cd "$pkgname"

  patch -p1 -i "$srcdir/remove-version-constraints.patch"

  # Fix version
  sed -e "s|0.0.0|$pkgver|" -i dropbox/dropbox_client.py
}

build() {
  cd "$pkgname"

  export SETUPTOOLS_SCM_PRETEND_VERSION=$pkgver

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  # temporary install
  python -m installer --destdir="$(pwd)/tmp" dist/*.whl
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  export PYTHONPATH="$(pwd)/tmp/$site_packages"

  pytest -vv --ignore test/integration/
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
