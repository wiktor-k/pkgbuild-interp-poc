# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: Dian M Fay <dian.m.fay@gmail.com>

pkgname=python-magic-wormhole-mailbox-server
pkgver=0.8.0
pkgrel=1
pkgdesc="the rendezvous/mailbox server for magic-wormhole clients "
arch=(any)
url="https://github.com/magic-wormhole/magic-wormhole-mailbox-server"
license=(MIT)
depends=(python-autobahn
         python-twisted)
makedepends=(python-{build,installer,wheel}
             python-setuptools
             python-versioneer)
checkdepends=(python-pytest
              python-treq)
source=(${pkgname}-${pkgver}.tar.gz::"${url}/archive/refs/tags/${pkgver}.tar.gz")
sha512sums=('2ee48c6d7d98a6dad9107cda6dd59b2ef7865436dce29433c3cd790c311527e807a9b4e8de879a1bb990521afec7a47077e1f0cc9006f7e7389597072bfd90d8')
b2sums=('5b67f9a8d1389b6735f29a627b16b0e41899c979abde6b4fdd637c87ffaa8ca312a8b4b1ba601a4ce1dd6f2b74d6cc4763974f8c2d5dac2204a58721536ea997')

prepare() {
	cd "${pkgname#python-}-$pkgver"
	# vendored versioneer is too old for Python 3.12
	rm -f versioneer.py
}

build() {
	cd "${pkgname#python-}-$pkgver"
	python -m build -wn
}

check() {
	cd "${pkgname#python-}-$pkgver"
	export PYTHONPATH='src'
	pytest
}

package() {
	cd "${pkgname#python-}-$pkgver"
	python -m installer -d "$pkgdir" dist/*.whl
	install -Dm0644 -t "$pkgdir/usr/share/licenses/$pkgname/" LICENSE
	install -Dm0644 -t "$pkgdir/usr/share/doc/$pkgname" docs/*.md
}
