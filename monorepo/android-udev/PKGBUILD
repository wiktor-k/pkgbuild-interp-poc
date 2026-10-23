# Maintainer: Robin Candau <antiz@archlinux.org>
# Contributor: M0Rf30
# Contributor: marlock

pkgname=android-udev
pkgver=20260922
pkgrel=1
pkgdesc="Udev rules to connect Android devices to your linux box"
url="https://github.com/M0Rf30/android-udev-rules"
arch=('any')
license=('GPL-3.0-only')
depends=('udev' 'libmtp')
source=("${pkgname}-${pkgver}.tar.gz::${url}/archive/${pkgver}.tar.gz")
sha256sums=('3cc14f34142e71837eb073cc10b09eb48a9f7282b7c9ade9259e7809eec5062e')

package() {
	cd "${pkgname}-rules-${pkgver}"
	install -Dm 644 51-android.rules "${pkgdir}/usr/lib/udev/rules.d/51-android.rules"
	install -Dm 644 android-udev.conf "${pkgdir}/usr/lib/sysusers.d/android-udev.conf"
}
