# Maintainer: kpcyrd <kpcyrd[at]archlinux[dot]org>
# Contributor: Felix Yan <felixonmars@archlinux.org>
# Contributor: T.J. Townsend <blakkheim@archlinux.org>

pkgname=minisign
pkgver=0.12
pkgrel=2
pkgdesc="A dead-simple tool to sign files and verify digital signatures"
arch=('x86_64')
url="https://github.com/jedisct1/minisign"
license=('ISC')
depends=(
  'glibc'
)
makedepends=(
  'git'
  'zig'
)
source=("git+https://github.com/jedisct1/minisign.git?signed#tag=$pkgver"
        zig-0.14.patch::https://github.com/jedisct1/minisign/commit/428c18e08a3fa704348e0c437747e4d4dd6612a9.diff)
sha512sums=('cd14bab2bdf384f410012b6f3a9149554fb367eafae187f11d02f4dd66b301a065e47215f95cd29150dce9b59a04d319b3b84c70bb555b60c8ebf6dd8bf25e55'
            'c19ea87ae57f2f1c950ee79644ebd9211c7bb23f1951240fe9a59ca60c2e8f86e67a90e6c1242567e791cc6410e036954d126cebe1679ede13175b76e7b89a3d')
validpgpkeys=(54A2B8892CC3D6A597B92B6C210627AABA709FE1) # Frank Denis

prepare() {
  cd "$pkgname"
  patch -Np1 -i ../zig-0.14.patch
}

build() {
  cd "$pkgname"
  zig build --verbose -Dcpu=baseline -Doptimize=ReleaseSmall -Dwithout-libsodium
}

check() {
  cd "$pkgname"
  zig-out/bin/minisign -GW -p test.pub -s test.sec
  echo 'ohai!' > msg.txt
  zig-out/bin/minisign -S -s test.sec -m msg.txt
  zig-out/bin/minisign -V -p test.pub -m msg.txt -x msg.txt.minisig
}

package() {
  cd "$pkgname"
  install -Dm755 -t "$pkgdir/usr/bin" zig-out/bin/minisign
  install -Dm644 -t "$pkgdir/usr/share/man/man1" share/man/man1/*
  install -Dm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
