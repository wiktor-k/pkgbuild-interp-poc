# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Ronald van Haren <ronald@archlinux.org>

pkgname=ttf-inconsolata
pkgver=3.000
pkgrel=5
epoch=1
pkgdesc='Monospace font for pretty code listings and for the terminal'
arch=(any)
url='https://fonts.google.com/specimen/Inconsolata'
license=(OFL-1.1)
source=(
  "$pkgname-$pkgver-fonts_ttf.zip::https://github.com/googlefonts/Inconsolata/releases/download/v$pkgver/fonts_ttf.zip"
  "$pkgname-$pkgver-OFL.txt::https://github.com/googlefonts/Inconsolata/releases/download/v$pkgver/OFL.txt"
  "$pkgname-$pkgver-FONTLOG.txt::https://github.com/googlefonts/Inconsolata/releases/download/v$pkgver/FONTLOG.txt"
  '75-ttf-inconsolata.conf'
)
sha512sums=('b83d0205642776752d7038f19b22de64a96268822d9a9f3501c3634bb80c96f9139e75200877e267c3aa1dcab37a282cb40155502adfcbe212d10c2b08137fd1'
            'ef825dc1608aa10addfd3fa2289d003f35e7365981434043e33d04f13b0f4197f09df685f2d6c162cd59cfebe2c525246f376e963874e09143f0fb0b7ba54325'
            '0983ecf9087924a8f2f140a2e963e1cf1e0ba5a9eba8ae406bbc934f6fcf782ad301d4018198975c02dd3d11558e0fb0160846b84da65d195a6a69de07bdaa95'
            '25e631954f02e611df679fb24534b1425931fa484ef97b7f24612cf300929b0f2e5a1389109602ceaee60f7ca2ea845b1b5eb190c5061ac48a907fe65428d72e')
b2sums=('9bf43df35f936edb562ab8075bbc1a77e3ac2db6cb35153768c145e388676d0c2878117e63853d51ddec5d4731b2ce601ef993e561fa22d3801777fbfc343ad8'
        '2d02874de9dce95ebf44a3869ff1cd813ec05a798be157fa25ff6436a81d3c617a28daecce77a26d2769ed3c13afd8bf48434a0d55e99362f3c78fb51884b637'
        '33d32b4b371fe8721e0c9e41590ed007677405bb0e77ce7432c3010c6b2522e1ad6082fa7c806649bedcb65f9f00b1ebe8dc02dc837bce666eb1ff485f122727'
        '80f258707cce1542cdfdd9c504325eba75b8580b93ad88260eb55bfb3a84c29c34420e027eb1ae01f4e3ec3939f2810327edff0c7b9535f8656ed6cb0129aedc')

package() {
  # fonts
  install -vDm644 -t "$pkgdir/usr/share/fonts/TTF" fonts/ttf/*.ttf

  # font config
  install -vDm644 -t "$pkgdir/usr/share/fontconfig/conf.avail" 75-ttf-inconsolata.conf

  # license
  install -vDm644 "$pkgname-$pkgver-OFL.txt" "$pkgdir/usr/share/licenses/$pkgname/LICENSE"

  # documentation
  install -vDm644 "$pkgname-$pkgver-FONTLOG.txt" "$pkgdir/usr/share/doc/$pkgname/FONTLOG.txt"
}
