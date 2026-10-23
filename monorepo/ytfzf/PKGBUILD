# Maintainer: Brett Cornwall <ainola@archlinux.org>
# Contributor: Gennadiy Mykhailiuta
# Contributor: Westofer Raymond

pkgname=ytfzf
pkgver=2.6.2
pkgrel=2
pkgdesc="A POSIX script to find and watch youtube videos from the terminal"
arch=('any')
url="https://github.com/pystardust/ytfzf"
license=('GPL-3.0-only')
depends=('jq' 'curl')
optdepends=(
    'dmenu: Search prompts and results via dmenu'
    'mpv: Default media player'
    'fzf: The default menu selection screen'
    'ueberzug: The default thumbnails image previewer'
    'yt-dlp: For downloading content'
)
source=("$pkgname-$pkgver.tar.gz::https://github.com/pystardust/ytfzf/archive/refs/tags/v$pkgver.tar.gz")
sha256sums=('73280e4ef4f490400a42bf582b713803c523587b0b30269859eaa6f2693ec9b1')

package() {
    cd "$pkgname-$pkgver"
    install -Dm0755 ytfzf -t "$pkgdir/usr/bin/"

    install -dm0755 "$pkgdir/usr/share/$pkgname/addons"
    cp -r addons/* "$pkgdir/usr/share/$pkgname/addons"
    find "$pkgdir/usr/share/$pkgname/addons" -exec chmod 0755 {} +

    install -Dm644 README.md -t "$pkgdir/usr/share/doc/$pkgname"
    cd docs
    install -Dm0644 "man/$pkgname.1" -t "$pkgdir/usr/share/man/man1"
    install -Dm0644 "man/$pkgname.5" -t "$pkgdir/usr/share/man/man5"
    install -Dm0644 conf.sh -t "$pkgdir/usr/share/doc/$pkgname"
    install -Dm0644 subscriptions -t "$pkgdir/usr/share/doc/$pkgname"
}
