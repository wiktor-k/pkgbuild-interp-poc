# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Contributor: Limao Luo <luolimao+AUR@gmail.com>
# Contributor: Wes Brown <wesbrown18@gmail.com>

pkgname=volatility3
pkgver=2.28.2
pkgrel=1
pkgdesc='Advanced memory forensics framework'
url='https://github.com/volatilityfoundation/volatility3'
arch=(any)
license=(LicenseRef-VolatilityFoundation)
depends=(
  python
  python-pefile
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-pytest
  python-yara
  unzip
)
optdepends=(
  'python-capstone: disassembly support'
  'python-pillow: screenshot and image support'
  'python-pycryptodome: cryptographic operations'
  'python-yara: YARA scanning support'
  'snappy: AVML compressed memory dump support'
)
provides=(volatility)
replaces=(volatility)
source=(
  "$url/archive/v$pkgver/$pkgname-$pkgver.tar.gz"
  "$pkgname-linux-sample-1.bin.gz::https://downloads.volatilityfoundation.org/volatility3/images/linux-sample-1.bin.gz"
  "$pkgname-symbols-linux.zip::https://downloads.volatilityfoundation.org/volatility3/symbols/linux.zip"
)
sha512sums=('f573bd89c59bbbab13837525db09c11eda4dc362fa367af231008c8b45b8e7276fc74b5fe0dfbc4889fc747e97b9682b415cd79d2a3397588a21b85e03400e8c'
            '82a34aa1d4e7d7deb9ea12e892f14ba0a28908b2da2e31ca76efd1e9e59d87769064a1ffffe8979a58b992cea7005fa20954df50984dbe1dd186513ad95168d7'
            '59056044c3702c1ffd5f1490cb457831d65d1f48fbc523da92b950963155df994a2ff2f11abd5f3b1a977c7f28ed48b3b256c88f81dfa96a66419cb50934ba49')
b2sums=('7edb9adf5de64379e139774aa7addb46c6ad1cbb38290029731a67de5cdd0a239fdfe1e356d2e1e1ab7678b1a7980f5bbfc8beb227839648be978dbb7981d7fb'
        '3321e991a50b6e4ccf19e0ee48d779664f43f4cdbdb950ca31a5d08c8ee9de0018e2d4f69f42206f7e8f1c6ea735c47f8aa42806ed0f85e9a837b611f07b3289'
        '53cd49663938c6f39e71193ae225e4985f8d0e6d301f2420eaa021c2722bfa6f9c9e7d029854b81f4cdd47bf2870e28efeff432fdc1281842b54ee9ffe6e3807')

build() {
  cd $pkgname-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd $pkgname-$pkgver
  mkdir -p test_images
  gzip -dc ../$pkgname-linux-sample-1.bin.gz > test_images/linux-sample-1.bin

  mkdir -p volatility3/symbols
  unzip ../$pkgname-symbols-linux.zip -d volatility3/symbols

  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m pytest ./test/plugins/linux/linux.py --volatility=volshell.py --image-dir=./test_images -k test_linux_volshell -v
  test-env/bin/python -m pytest ./test/plugins/linux/linux.py --volatility=vol.py --image-dir=./test_images -k "test_linux and not test_linux_volshell" -v --durations=0
}

package() {
  cd $pkgname-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 -t "$pkgdir/usr/share/doc/$pkgname" README.md
  install -vDm 644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE.txt
}
