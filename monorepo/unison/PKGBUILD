# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: xs5871 <xs5871@qn-4.net>
# Contributor: Gaetan Bisson <bisson@archlinux.org>
# Contributor: Tobias Powalowski <tpowa@archlinux.org>

pkgname=unison
pkgver=2.54.0
pkgrel=1
pkgdesc='File-synchronization tool'
url="https://github.com/bcpierce00/$pkgname/"
arch=(x86_64)
license=(GPL-3.0-only)
optdepends=('gtk3: graphical interface')
makedepends=(lablgtk3
             ocaml)
_archive="$pkgname-$pkgver"
source=("$url/archive/v$pkgver/$_archive.tar.gz"
        "$pkgname.desktop")
sha512sums=('7d1fea12b11e7a4bf2a5dbd9064d3c5104ca21df88fae0ead098f1ec8b7b935abcc7cdff74799fd2d9f5b23cc06eaa7f262bb71270511fb9c2ff4aac74c7ff11'
            '3e6e7b793deb60dab64c6e585e3b2ceee01ced0f8b73d1373ac1cd8aaaadb5c7415f2295b8a044a57fc15356dfe76bab002e1e633495e5f8e436871c83c48990')
options=('!makeflags')

build() {
	cd "$_archive"
	make all
}

package() {
	cd "$_archive"
	install -Dm0755 -t "$pkgdir/usr/bin/" "src/$pkgname"{,-fsmonitor,-gui}
	install -Dm0644 -t "$pkgdir/usr/share/applications/" "../$pkgname.desktop"
	install -Dm0644 -t "$pkgdir/usr/share/man/man1/" man/*.1
	install -Dm0644 icons/U.svg "$pkgdir/usr/share/icons/hicolor/scalable/apps/$pkgname.svg"
	for i in 16 24 32 48 256; do
		install -Dm0644 "icons/U.${i}x${i}x16m.png" "$pkgdir/usr/share/icons/hicolor/${i}x${i}/apps/$pkgname.png"
	done
}
