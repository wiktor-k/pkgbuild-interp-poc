# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>
# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: Storm Dragon <stormdragon2976@gmail.com>

pkgname=magic-wormhole
pkgver=0.24.0
pkgrel=1
pkgdesc="Securely transfer data between computers"
arch=(any)
url="https://github.com/magic-wormhole/magic-wormhole"
license=(MIT)
depends=(python-autobahn
         python-click
         python-humanize
         python-idna
         python-iterable-io
         python-pynacl
         python-pyopenssl
         python-qrcode
         python-service-identity
         python-spake2
         python-tqdm
         python-txtorcon
         python-zipstream-ng)
makedepends=(python-{build,installer,wheel}
             python-setuptools
             python-versioneer)
checkdepends=(python-magic-wormhole-mailbox-server
              python-magic-wormhole-transit-relay
              python-noiseprotocol
              python-pytest
              python-pytest-mock
              python-pytest-twisted
              python-hypothesis)
_archive="$pkgname-$pkgver"
source=("$url/archive/$pkgver/$_archive.tar.gz")
sha512sums=('fb4c4430456a20575630dbdddea5557ef050e974da0306d5e9c6d2eedec43417444c7e17c7e6416fb6118dd6007a8fd420f40351fb358c107842e682ae2a2e0e')

prepare() {
	cd "${pkgname#python-}-$pkgver"
	# vendored versioneer is too old for Python 3.12
	rm -f versioneer.py
}

build() {
	cd "$_archive"
	python -m build -wn
}

check() {
	cd "$_archive"
	export PYTHONPATH=src
	local deselected=(
		# uses unittest.mock wrong and triggers Python errors
		src/wormhole/test/test_tor_manager.py
		# use network connections via twisted with flakey race conditions
		src/wormhole/test/test_args.py
		src/wormhole/test/test_cli.py
	)
	pytest ${deselected[@]/#/--deselect }
}

package() {
	cd "$_archive"
	python -m installer -d "$pkgdir" dist/*.whl
	install -vDm 644 -t "$pkgdir/usr/share/licenses/$pkgname/" LICENSE
	install -vDm 644 -t "$pkgdir/usr/share/man/man1/" docs/wormhole.1
	mkdir -p "$pkgdir/usr/share/"{bash-completion/completions,fish/vendor_completions.d,zsh/site-functions}/
	mv "$pkgdir/usr/"{wormhole_complete.bash,share/bash-completion/completions/wormhole}
	mv "$pkgdir/usr/"{wormhole_complete.fish,share/fish/vendor_completions.d/wormhole.fish}
	mv "$pkgdir/usr/"{wormhole_complete.zsh,share/zsh/site-functions/_wormhole}
}
