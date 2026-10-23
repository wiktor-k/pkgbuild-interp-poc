# Maintainer: David Runge <dvzrv@archlinux.org>
# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Maintainer: Christian Heusel <gromit@archlinux.org>

_name=pynitrokey
pkgname=python-pynitrokey
pkgver=0.14.1
pkgrel=1
pkgdesc="A command line interface for the Nitrokey FIDO2, Nitrokey Start, Nitrokey 3 and NetHSM"
arch=(any)
url="https://github.com/Nitrokey/pynitrokey"
license=(
  'Apache-2.0 OR MIT'
  LGPL-3.0-only
  GPL-3.0-or-later
)
depends=(
  nitrokey-udev-rules
  python
  python-cffi
  python-click
  python-cryptography
  python-fido2
  python-intelhex
  python-libusb1
  python-nethsm
  python-nitrokey
  python-nkdfu
  python-pyserial  # implicit via python-nitrokey
  python-pyusb
  python-requests
  python-semver
  python-tlv8
  python-tqdm
)
makedepends=(
  git
  python-build
  python-poetry-core
  python-installer
  python-wheel
)
# NOTE: there are no tests to run
optdepends=(
  'python-pyscard: for PC/SC support'
)
source=(
  "$_name::git+$url.git?signed#tag=v$pkgver"
)
sha512sums=('2ebde2d498717461f3c80e817bd7ab577754bfc79274de348c1b73f755a41f9fa8a857a2abbcf07e85bbff99b9234d64672cbc810b9dbf078998a713ccf2f6e0')
b2sums=('9eb2634eed87e337a739995d3896998842f051b8656fd6cc402623cba579d581423376d08ccaa36e5980b28b6faab3276a5e73961b21eea7f15431d6ced4bfc1')
validpgpkeys=(
  868184069239FF65DE0BCD7DD9BAE35991DE5B22  # Szczepan Zalega <szczepan@nitrokey.com> (@szszszsz)
  CC74B7120BFAA36FF42868724C1449F1C9804176  # Markus Meissner <meissner@nitrokey.com> (@daringer)
  97591008D82D0751EF438D35804551D1A83A3C83  # Markus Merklinger <markus@nitrokey.com> (@mmerklinger)
  463D8C09EC635FC7DA5681439C44D117B05429AC  # Robin Krahl <robin@nitrokey.com> (@robin-nitrokey)
  93CCB0DB717DAE30622F671436DA48A4C827B354  # Sosthène Guédon <sosthene@nitrokey.com> (@sosthene-nitrokey)
)

build() {
  cd $_name
  python -m build --wheel --skip-dependency-check --no-isolation
}

package() {
  cd $_name
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 README.md -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 LICENSES/*.txt -t "$pkgdir/usr/share/licenses/$pkgname/"
}
