# Maintainer: Justin Kromlinger <hashworks@archlinux.org>
# Contributor: Chih-Hsuan Yen <yan12125@archlinux.org>
# Contributor: Abdó Roig-Maranges <abdo.roig@gmail.com>

_pkgname=jwcrypto
pkgname=python-$_pkgname
pkgver=1.6.1
pkgrel=1
pkgdesc='Python implementation of JWK, JWS, JWE specifications'
arch=(any)
url='https://github.com/latchset/jwcrypto'
# setup.py says LGPLv3+ https://github.com/latchset/jwcrypto/blob/v1.5.1/setup.py#L20
license=(LGPL-3.0-or-later)
# This package uses @typing_extensions.deprecated, which will be available as @warnings.deprecated in Python 3.13
# https://github.com/latchset/jwcrypto/pull/337
# https://typing-extensions.readthedocs.io/en/latest/#deprecated
depends=(
  python
  python-cryptography
  python-typing_extensions
)
makedepends=(
  git
  python-build
  python-installer
  python-hatchling
)
checkdepends=(python-pytest)
source=("git+$url?signed#tag=v${pkgver}")
sha512sums=('992dac2b6b360e82fed8a5061f66e96538db4a0cb10588d686c9be6c9ff030ea0534b684dc89cf3ad3a6edcea615ad4741d546e603cbea3e096896a949006c67')
b2sums=('1864436f3dcf7dd97c353ee852cdc870dd41abc8d01cb71e49765a4176d3b33617cb2c5dc48bc4461529095db3535e8f1881f0547c18f10911ebc9e69c0945a9')
# PyPI sdists for older versions are signed by tiran's key. Since version 1.0,
# PyPI sdists are not signed, and tiran didn't declare transition of the
# signing key for some reason. I temporarily trust simo5's signed git tags as
# the @redhat.com UID in their key has been verified on keys.openpgp.org.
# See https://github.com/latchset/jwcrypto/issues/230 for more details.
validpgpkeys=(
  'BB97AF8BC4E7A5C0D96223D3C788C4C1D4550D45'  # https://github.com/tiran
  '7C7BD146943B206BB645B64594EAD67E004B65AB'  # Simo Sorce <simo@redhat.com>
)

build() {
  cd $_pkgname
  python -m build --wheel --no-isolation
}

check() {
  cd $_pkgname
  pytest
}

package() {
  cd $_pkgname
  python -m installer --destdir="$pkgdir" dist/*.whl
}
