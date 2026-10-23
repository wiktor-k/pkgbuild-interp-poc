# Maintainer: BlackIkeEagle <ike DOT devolder AT gmail DOT com>
# Contributor: Sergej Pupykin <pupykin.s+arch@gmail.com>
# Contributor: John Gerritse <john.gerritse@gmail.com>

pkgname=atop
pkgver=2.13.0
pkgrel=1
pkgdesc='A system and process level monitor'
arch=('x86_64')
url='https://www.atoptool.nl/'
license=('GPL-2.0-or-later')
depends=('ncurses' 'sh' 'zlib')
optdepends=(
    'python-pynvml: NVIDIA gpu reporting (AUR)'
)
backup=('etc/default/atop')
source=("https://www.atoptool.nl/download/atop-${pkgver/_/-}.tar.gz")
sha512sums=('646e5c2d224fafd7635395cef6b5ad173dda872b8c39c120e0e31b9634a4b04174b7e5ce6885fc6c35e72c11a41f450a2c29fc648512b81f11cf55532b19220d')

#prepare() {
    #cd "$pkgname-$pkgver"
    #patch -p1 -i "$srcdir/262.patch"
#}

build() {
    cd "$pkgname-$pkgver"
    make
}

package() {
    cd "$pkgname-$pkgver"
    ## default folder not properly created
    mkdir -p "$pkgdir/etc/default"
    make DESTDIR="$pkgdir" install \
        SBINPATH=/usr/bin \
        SYSDPATH=/usr/lib/systemd/system
    install -Dm644 "atop.default" "$pkgdir/etc/default/atop"
    # enable atop-rotate.timer by default
    install -dm755 "$pkgdir/usr/lib/systemd/system/timers.target.wants"
    ln -s ../atop-rotate.timer \
        "$pkgdir/usr/lib/systemd/system/timers.target.wants/atop-rotate.timer"
}
