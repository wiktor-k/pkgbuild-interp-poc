# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-softlayer
pkgver=6.2.9
pkgrel=1
pkgdesc="A library for SoftLayer's API"
arch=('any')
url='https://github.com/softlayer/softlayer-python'
license=('MIT')
depends=('python' 'python-click' 'python-requests'
         'python-prompt_toolkit' 'python-urllib3' 'python-rich')
makedepends=('git' 'python-build' 'python-installer' 'python-setuptools' 'python-wheel')
checkdepends=('python-pytest' 'python-testtools' 'tk')
source=("git+https://github.com/softlayer/softlayer-python.git#tag=v$pkgver"
         "$pkgname-click-8.3-compatibility.patch")
sha512sums=('2a5dd5076f9a8e50e35a65ec45424060203230bbe542ee360ad55673dd0e4b8aae7fe79174ee4d9ab696fd1d09abfde2bd5fb21171800e653b8d0c96d3ddc119'
            'f2be77d7dda539a9fc8eb63ffbee7f0bba3e1e1bf72768bff674169158502f1f2bf36c44f99e20e98d2e1cc377f1c6a01805a78133ae087b9aa5f8eca0e8b829')

prepare() {
  cd softlayer-python
  sed -i 's/==/>=/' setup.py
  patch -p1 -i ../$pkgname-click-8.3-compatibility.patch
  # https://github.com/softlayer/softlayer-python/commit/681d581b fixes nondeterministic cf_call test ordering.
  git show 681d581b6f15e490f2c9bfc971f2fae7f0dad8c2 -- tests/api_tests.py | git apply
}

build() {
  cd softlayer-python
  python -m build --wheel --no-isolation
}

check() {
  cd softlayer-python
  pytest -v
}

package() {
  cd softlayer-python
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/

  # It's only present when tests are enabled, so adding -f
  local python_version=$(python -c 'import sys; print(".".join(map(str, sys.version_info[:2])))')
  rm -rf "$pkgdir"/usr/lib/python${python_version}/site-packages/tests

  # Prevent collisions with `sl` package; This command is deemed deprecated by
  # the project anyway
  rm -f "$pkgdir/usr/bin/sl"
}

# vim:set ts=2 sw=2 et:
