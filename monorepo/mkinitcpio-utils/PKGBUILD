# Maintainer: Giancarlo Razzolini <grazzolini@archlinux.org>
pkgname=mkinitcpio-utils
pkgver=0.0.6
pkgrel=1
pkgdesc="Collection of Archlinux mkinitcpio utilities performing various tasks"
arch=('any')
url="https://github.com/grazzolini/mkinitcpio-utils"
license=('BSD-3-Clause')
depends=('cryptsetup')
optdepends=('mkinitcpio-dropbear: Allow the encryptssh hook to unlock a root container remotely using dropbear'
            'mkinitcpio-tinyssh: Allow the encryptssh hook to unlock a root container remotely using tinyssh')
changelog='ChangeLog'
source=("${pkgname}-${pkgver}.tar.gz::$url/archive/v$pkgver.tar.gz")
sha512sums=('cdc82fc034a781e603716540b34ec325c3835ce0ad797f3f352bbe87165325bdfed3129bbc10187d8592be162b63350aba774e5664de78bd9e38639cc95dbfff')

package() {
  install -Dm755 "$srcdir/$pkgname-$pkgver/utils/shells/cryptsetup_shell"       "$pkgdir/usr/share/$pkgname/utils/shells/cryptsetup_shell"
  install -Dm644 "$srcdir/$pkgname-$pkgver/README.md"                           "$pkgdir/usr/share/$pkgname/README.md"
  install -Dm644 "$srcdir/$pkgname-$pkgver/initcpio/hooks/encryptssh"           "$pkgdir/usr/lib/initcpio/hooks/encryptssh"
  install -Dm644 "$srcdir/$pkgname-$pkgver/initcpio/install/encryptssh"         "$pkgdir/usr/lib/initcpio/install/encryptssh"
  install -Dm644 "$srcdir/$pkgname-$pkgver/LICENSE"                             "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
}
