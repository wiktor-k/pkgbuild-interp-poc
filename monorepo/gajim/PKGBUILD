# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Maintainer: Robin Candau <antiz@archlinux.org>
# Contributor: Eric Bélanger <eric@archlinux.org>

pkgname=gajim
pkgver=2.6.0
pkgrel=4
pkgdesc="Full featured and easy to use XMPP (Jabber) client"
url="https://gajim.org/"
arch=('any')
license=('GPL-3.0-only')
depends=('gtk4' 'gtksourceview5' 'gstreamer' 'libadwaita' 'python-cairo' 'python-gobject' 'python-keyring' 'python-nbxmpp' 'python-cryptography' 'python-precis_i18n' 'python-css-parser' 'python-distro' 'hicolor-icon-theme' 'python-pillow' 'python-gssapi' 'python-netifaces' 'python-qrcode' 'python-omemo-dr' 'python-packaging' 'pango' 'sqlite' 'python-sqlalchemy' 'python-emoji' 'python-httpx' 'python-truststore' 'python-h2' 'python-socksio' 'emoji-font' 'python-pysequoia')
makedepends=('python-build' 'python-installer' 'python-wheel' 'python-setuptools')
optdepends=('python-dbus: to have gajim-remote working'
            'python-sentry_sdk: for Sentry error reporting (users decide whether to send reports or not)'
            'libspelling: for spell checking support'
            'libsecret: for GNOME Keyring or KDE support as password storage'
            'gupnp-igd: for better NAT traversing'
            'networkmanager: for network lose detection'
            'geoclue: share current location'
            'gsound: Notification Sounds'
            'farstream: for video and audio calls'
            'gst-plugins-base: for video and audio calls'
            'gst-plugins-ugly: for video and audio calls'
            'gst-libav: for video and audio calls'
            'gst-plugin-gtk4: for video and audio calls'
            'libxss: for idle time checking on X11'
            'python-gpgme: encrypting chat messages with OpenPGP')
source=("https://gitlab.com/gajim/gajim/-/archive/${pkgver}/gajim-${pkgver}.tar.gz")
sha512sums=('8cc3678da8c7d73b9bce58afc242cf71608b59ac210fd170d128868505d95a3f9788fb69571e476c0ad325f3bc0476016a9cc6670604c72317b7dbd4f0a40f4e')
b2sums=('0a7caf9b95c1bf7f739cfa43fd70d08951b59f46a9c7ed54198c6cf6778d3ce59265ff8d62360f76d5db612ce5dac651b0f4d4e51d0cb86531b6773909dccbbc')


build() {
	cd "${pkgname}-${pkgver}"
	./make.py build --dist unix
	python -m build --wheel --no-isolation
}

check() {
	cd "${pkgname}-${pkgver}"
	python -m unittest discover -s test
}

package() {
	cd "${pkgname}-${pkgver}"
	python -m installer --destdir="${pkgdir}" dist/*.whl
	./make.py install --dist unix --prefix="${pkgdir}/usr"
}
