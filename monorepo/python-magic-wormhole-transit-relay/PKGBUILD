# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Maintainer: Caleb Maclennan <caleb@alerque.com>

pkgname=python-magic-wormhole-transit-relay
pkgver=0.5.0
pkgrel=1
pkgdesc="transit Relay server for Magic-Wormhole"
arch=(any)
url="https://github.com/magic-wormhole/magic-wormhole-transit-relay"
license=(MIT)
depends=(python-autobahn
         python-twisted)
makedepends=(python-{build,installer,wheel}
             python-setuptools
             python-versioneer)
checkdepends=(python-pytest)
source=(${pkgname}-${pkgver}.tar.gz::"https://github.com/magic-wormhole/magic-wormhole-transit-relay/archive/refs/tags/${pkgver}.tar.gz")
sha512sums=('13b2358ab7c076bb822cd9e2394ecdd6d30769a043ed2e2cf04b68f451cd28dbedbfbe329f945581e9c0a83b2bde7b1a2529d8d38ef29a66088190aedbcd7d1f')
b2sums=('7c2e98de1042410558cd2a7e219f821c71fa7e90b84b7cae1498dc155baf36aeef4f0a907f9c2419175bbae0d28b8826c4d8f1a8ee41088b8dad36e21d884963')

prepare() {
	cd "${pkgname#python-}-$pkgver"
	# De-vendor old versioneer copy, incompatible with 3.12
	rm versioneer.py
}

build() {
	cd "${pkgname#python-}-${pkgver}"
	python -m build -wn
}

check() {
	cd "${pkgname#python-}-${pkgver}"
	export PYTHONPATH='src/wormhole_transit_relay'
	pytest --deselect src/wormhole_transit_relay/test/test_backpressure.py
}

package() {
	cd "${pkgname#python-}-${pkgver}"
	python -m installer -d "$pkgdir" dist/*.whl
	install -Dm0644 -t "$pkgdir/usr/share/licenses/$pkgname/" LICENSE
	install -Dm0644 -t "$pkgdir/usr/share/doc/$pkgname" docs/*.md
}
